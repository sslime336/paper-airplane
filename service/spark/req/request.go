package req

import "sync"

type Request struct {
	Header    Header    `json:"header"`
	Parameter Parameter `json:"parameter"`
	Payload   Payload   `json:"payload"`
}

type Header struct {
	AppID string `json:"app_id"`
	Uid   string `json:"uid,omitempty"`
}

type Parameter struct {
	Chat Chat `json:"chat"`
}

type Chat struct {
	// Domain 模型选择
	// 取值为[general,generalv2,generalv3,pro-128k,generalv3.5,4.0Ultra]
	Domain string `json:"domain"`

	// Temperature 核采样阈值。用于决定结果随机性，取值越高随机性越强即相同的问题得到的不同答案的可能性越高
	// 取值范围 (0，1] ，默认值0.5
	Temperature float64 `json:"temperature,omitempty"`

	// MaxTokens 模型回答的tokens的最大长度, Lite、Pro、Max、4.0 Ultra 取值为[1,8192]，默认为4096
	MaxTokens int64 `json:"max_tokens,omitempty"`

	// TopK 从k个候选中随机选择⼀个（⾮等概率）, 取值为[1，6],默认为4
	TopK     int64  `json:"top_k,omitempty"`
	Auditing string `json:"auditing"`
}

type Payload struct {
	Message Message `json:"message"`
}

type Message struct {
	Text []Text `json:"text"`
}

func (msg *Message) Add(role, content string) {
	msg.Text = append(msg.Text, Text{Role: role, Content: content})
}

type Text struct {
	// Role system用于设置对话背景，user表示是用户的问题，assistant表示AI的回复
	// 取值为[system,user,assistant]
	Role string `json:"role"`

	// Content 用户和AI的对话内容
	// Lite、Pro、Max、4.0 Ultra版本: 所有content的累计tokens需控制8192以内
	Content string `json:"content"`
}

var sparkLiteRequestPool = &sync.Pool{
	New: func() any {
		return &Request{
			Header: Header{
				AppID: "",
			},
			Parameter: Parameter{
				Chat: Chat{
					Domain:      "general",
					Temperature: 0.8,
					MaxTokens:   2048,
					TopK:        6,
					Auditing:    "default",
				},
			},
			Payload: Payload{
				Message: Message{
					Text: []Text{},
				},
			},
		}
	},
}

// NewSparkLiteRequest 从对象池中拉取一个基于 SparkLite 模型的请求对象
func NewSparkLiteRequest(appId string) *Request {
	req := sparkLiteRequestPool.Get().(*Request)
	req.Header.AppID = appId
	return req
}
