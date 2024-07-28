package resp

type Response struct {
	Header  Header  `json:"header"`
	Payload Payload `json:"payload"`
}

type Header struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`

	// Sid 会话的唯一id，用于讯飞技术人员查询服务端会话日志使用,出现调用错误时建议留存该字段
	Sid string `json:"sid"`

	// Status 会话状态，取值为[0,1,2]；0代表首次结果；1代表中间结果；2代表最后一个结果
	Status int64 `json:"status"`
}

type Payload struct {
	Choices Choices `json:"choices"`
	Usage   Usage   `json:"usage"`
}

type Choices struct {
	// Status 	文本响应状态，取值为[0,1,2]; 0代表首个文本结果；1代表中间文本结果；2代表最后一个文本结果
	Status int64 `json:"status"`

	// Seq 返回的数据序号，取值为[0,9999999]
	Seq  int64         `json:"seq"`
	Text []TextElement `json:"text"`
}

type TextElement struct {
	// Content AI的回答内容
	Content string `json:"content"`
	// Role 角色标识，固定为assistant，标识角色为AI
	Role string `json:"role"`
	// Index 结果序号，取值为[0,10]; 当前为保留字段，开发者可忽略
	Index int64 `json:"index,omitempty"`
}

type Usage struct {
	Text UsageText `json:"text"`
}

type UsageText struct {
	// QuestionTokens 保留字段
	QuestionTokens int64 `json:"question_tokens,omitempty"`
	// PromptTokens 包含历史问题的总tokens大小，不同模型需要检测是否超出
	PromptTokens int64 `json:"prompt_tokens"`
	// CompletionTokens 回答的tokens大小
	CompletionTokens int64 `json:"completion_tokens"`
	// TotalTokens prompt_tokens和completion_tokens的和，也是本次交互计费的tokens大小
	TotalTokens int64 `json:"total_tokens"`
}
