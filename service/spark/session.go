package spark

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"paper-airplane/config"
	"paper-airplane/db"
	"paper-airplane/logging"
	"paper-airplane/service/spark/req"
	"paper-airplane/service/spark/resp"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// Session 代表一次与讯飞模型的会话
type Session struct {
	*websocket.Conn

	// SystemContent 对话背景
	// e.g. "你现在扮演李白，你豪情万丈，狂放不羁；接下来请用李白的口吻和用户对话。"
	SystemContent string

	// Message 包括所有的历史会话和新加入的会话
	Message req.Message
}

var ansPool = &sync.Pool{
	New: func() any {
		return &strings.Builder{}
	},
}

var log *zap.Logger

func Init() {
	log = logging.Logger().Named("spark.session")
	initCache()
}

var ws = websocket.Dialer{
	HandshakeTimeout: 5 * time.Second,
}

// Send 发送消息
func (s *Session) Send(msg string) error {
	r := req.NewSparkLiteRequest(config.App.Spark.AppId)

	var err error
	var resp *http.Response
	if s.Conn, resp, err = ws.Dial(buildAuthUrl(HostUrlSparkLite, config.App.Spark.ApiKey, config.App.Spark.ApiSecret), nil); err != nil {
		log.Error("create websocket connection failed", zap.String("response", readResp(resp)), zap.Error(err))
		return err
	}
	s.Message.Add(req.Text{
		Role:    "user",
		Content: msg,
	})
	r.Payload.Message.Text = append([]req.Text{}, s.Message.Text...)

	log.Debug("build spark lite request", zap.Any("request", *r))
	return s.WriteJSON(r)
}

// Read 阻塞等待获取回答
func (s *Session) Read() (string, error) {
	answer := ansPool.Get().(*strings.Builder)
	for {
		_, msg, err := s.ReadMessage()
		if err != nil {
			log.Error("read message error", zap.Error(err))
			return "", err
		}

		rsp := resp.Response{}
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

		currentContent := rsp.Payload.Choices.Text[0].Content

		answer.WriteString(currentContent)

		// AI 回答生成完毕
		if rsp.Header.Status == 2 {
			s.Close()
			break
		}
	}

	res := answer.String()

	s.Message.Add(req.Text{
		Role:    "assistant",
		Content: res,
	})

	answer.Reset()
	ansPool.Put(answer)

	return res, nil
}

// buildAuthUrl 创建需要的鉴权 URL
func buildAuthUrl(hosturl string, apiKey, apiSecret string) string {
	ul, err := url.Parse(hosturl)
	if err != nil {
		logging.Error("failed to parse Spark host url", zap.Error(err))
		return ""
	}
	date := time.Now().UTC().Format(time.RFC1123)

	signFields := []string{"host: " + ul.Host, "date: " + date, "GET " + ul.Path + " HTTP/1.1"}

	sign := strings.Join(signFields, "\n")
	logging.Debug("generate spark signature string", zap.String("sign", sign))

	sha := hmacWithShaTobase64(sign, apiSecret)

	logging.Debug("hmac with sha to base64", zap.String("base64", sha))

	authUrl := fmt.Sprintf("hmac username=\"%s\", algorithm=\"%s\", headers=\"%s\", signature=\"%s\"", apiKey, "hmac-sha256", "host date request-line", sha)

	authorization := base64.StdEncoding.EncodeToString([]byte(authUrl))

	v := url.Values{}
	v.Add("host", ul.Host)
	v.Add("date", date)
	v.Add("authorization", authorization)

	callurl := hosturl + "?" + v.Encode()
	return callurl
}

func hmacWithShaTobase64(data, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	encodeData := mac.Sum(nil)
	return base64.StdEncoding.EncodeToString(encodeData)
}

const HostUrlSparkLite = "wss://spark-api.xf-yun.com/v1.1/chat"

func NewSparkSession(openId string) (*Session, error) {
	var session *Session
	if sess, ok := sessionCache.Get(openId); ok {
		session = sess.(*Session)
	} else {
		var msess db.Session
		if err := db.Sqlite.Model(&db.Session{}).Find(&msess, "openId = ?", openId).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			session = &Session{}
		} else {
			session = &Session{
				SystemContent: "",
				Message:       req.Message{},
			}
			json.Unmarshal(msess.MessageHistory, &session.Message)
		}
		sessionCache.SetDefault(openId, session)
	}

	var resp *http.Response
	var err error
	session.Conn, resp, err = ws.Dial(buildAuthUrl(HostUrlSparkLite, config.App.Spark.ApiKey, config.App.Spark.ApiSecret), nil)
	if err != nil {
		log.Error("websocket to spark failed", zap.Error(err), zap.String("reason", readResp(resp)))
		return nil, err
	} else if resp.StatusCode != 101 {
		log.Error("websocket to spark failed, status code not 101", zap.Error(err), zap.String("reason", readResp(resp)))
		return nil, errors.New("unexpected status code")
	}

	return session, nil
}

func readResp(resp *http.Response) string {
	if resp == nil {
		return ""
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("code=%d,body=%s", resp.StatusCode, string(b))
}
