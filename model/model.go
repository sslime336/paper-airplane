package model

import "time"

type WsDtoBase struct {
	Op int64  `json:"op"`
	S  int64  `json:"s"`
	T  string `json:"t"`
	Id string `json:"id"`
}

type General struct {
	WsDtoBase
	D any `json:"d"`
}

// PrivateChatMessage 群聊中被 @
type PrivateChatMessage struct {
	WsDtoBase
	Data PrivateChatMessageD `json:"d"`
}

type PrivateChatMessageD struct {
	Author struct {
		Id         string `json:"id"`
		UserOpenid string `json:"user_openid"`
	} `json:"author"`
	Content string `json:"content"`
	ID      string `json:"id"`
}

type GroupAtMessage struct {
	WsDtoBase
	Data GroupAtMessageD `json:"d"`
}

type GroupAtMessageD struct {
	Author struct {
		ID           string `json:"id"`
		MemberOpenid string `json:"member_openid"`
	} `json:"author"`
	Content     string    `json:"content"`
	GroupID     string    `json:"group_id"`
	GroupOpenid string    `json:"group_openid"`
	ID          string    `json:"id"`
	Timestamp   time.Time `json:"timestamp"`
}
