package spark

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sslime336/paper-airplane/dao"
	"github.com/sslime336/paper-airplane/db/orm"
	"github.com/sslime336/paper-airplane/service/spark/req"
	"github.com/sslime336/paper-airplane/service/spark/resp"
	"go.uber.org/zap"
	"gorm.io/gen/field"
)

// Session 代表一次与讯飞模型的会话
type Session struct {
	*websocket.Conn `json:"-"`

	// SystemContent 对话背景
	// e.g. "你现在扮演李白，你豪情万丈，狂放不羁；接下来请用李白的口吻和用户对话。"
	SystemContent string

	// Message 包括所有的历史会话和新加入的会话
	Message req.Message

	TotalTokens int64
}

func (s *Session) JsonizedMessage() string {
	data, _ := json.Marshal(s.Message)
	return string(data)
}

// TokenOverflow 判断是否超出最大token数
func (s *Session) TokenOverflow() bool {
	return s.TotalTokens > 8192
}

// ResetToken 清空历史token数，并刷新会话历史
func (s *Session) ResetToken() {
	s.TotalTokens = 0
	s.Message = req.Message{}
}

var ws = websocket.Dialer{
	HandshakeTimeout: 5 * time.Second,
}

func NewSparkSession(openId string) (*Session, error) {
	var session *Session
	if sess, ok := sessionCache.Get(openId); ok {
		session = sess.(*Session)
	} else {
		session = &Session{}
		if daoSession, err := dao.Session.Attrs(field.Attrs(&orm.Session{OpenId: openId, JsonizedHistory: "", TotalTokens: 0})).Where(dao.Session.OpenId.Eq(openId)).FirstOrInit(); err != nil {
			return nil, err
		} else {
			session.TotalTokens = daoSession.TotalTokens
			json.Unmarshal([]byte(daoSession.JsonizedHistory), &session.Message)
			sessionCache.SetDefault(openId, session)
		}
	}

	var resp *http.Response
	var err error
	session.Conn, resp, err = ws.Dial(authUrl(), nil)
	if err != nil {
		log.Error("websocket to spark failed", zap.Error(err), zap.String("reason", readResp(resp)))
		return nil, err
	} else if resp.StatusCode != 101 {
		log.Error("protocol switch to websocket failed", zap.Error(err), zap.String("reason", readResp(resp)))
		return nil, errors.New("协议切换至 websocket 失败")
	}

	return session, nil
}

var createSparkLiteRequest func() *req.Request

// Send 发送消息
func (s *Session) Send(msg string) error {
	r := createSparkLiteRequest()

	var err error
	var resp *http.Response
	if s.Conn, resp, err = ws.Dial(authUrl(), nil); err != nil {
		log.Error("create websocket connection failed", zap.String("response", readResp(resp)), zap.Error(err))
		return err
	}

	if s.TokenOverflow() {
		s.ResetToken()
	}

	s.Message.Add("user", msg)

	r.Payload.Message.Text = append([]req.Text{}, s.Message.Text...)

	log.Debug("build spark lite request", zap.Any("request", *r))
	return s.WriteJSON(r)
}

// Read 阻塞等待获取回答
func (s *Session) Read() (string, error) {
	var answer strings.Builder
	for {
		_, msg, err := s.ReadMessage()
		if err != nil {
			log.Error("read message error", zap.Error(err))
			return "", err
		}

		var rsp resp.Response
		if err := json.Unmarshal(msg, &rsp); err != nil {
			log.Error("failed unmarshal response", zap.Error(err))
			return "", err
		}

		if rsp.Header.Code != 0 {
			log.Error("encountered error", zap.Int64("code", rsp.Header.Code))
			return "", ErrorCodeMap[rsp.Header.Code]
		}

		switch rsp.Header.Status {
		case 0:
			log.Debug("first session", zap.String("sid", rsp.Header.Sid))
		case 1:
			log.Debug("middle session", zap.String("sid", rsp.Header.Sid))
		case 2:
			log.Debug("final session", zap.String("sid", rsp.Header.Sid))
		}

		totalToken := rsp.Payload.Usage.Text.TotalTokens
		s.TotalTokens += totalToken

		// SparkLite 最多 Content token 数为 8192
		log.Debug("total token", zap.Int64("totalToken", rsp.Payload.Usage.Text.TotalTokens))

		currentContent := rsp.Payload.Choices.Text[0].Content

		answer.WriteString(currentContent)

		// AI 回答生成完毕
		if rsp.Header.Status == 2 {
			s.Close()
			break
		}
	}

	res := answer.String()

	s.Message.Add("assistant", res)

	return res, nil
}
