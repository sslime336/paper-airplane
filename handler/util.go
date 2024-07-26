package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"paper-airplane/logging"
	"paper-airplane/model"

	"go.uber.org/zap"
)

func SendTextMessagePrivate(userOpenId, prvMsgId, msg string) (err error) {
	url := fmt.Sprintf("https://sandbox.api.sgroup.qq.com/v2/users/%s/messages", userOpenId)
	_, err = api.Transport(context.Background(), "POST", url, SendMessage{
		Content: msg,
		MsgType: 0,
		MsgId:   prvMsgId,
	})
	return
}

func SendTextMessageGroup(groupId, prvMsgId, msg string) (err error) {
	url := fmt.Sprintf("https://sandbox.api.sgroup.qq.com/v2/groups/%s/messages", groupId)
	_, err = api.Transport(context.Background(), "POST", url, SendMessage{
		Content: msg,
		MsgType: 0,
		MsgId:   prvMsgId,
	})
	return
}

type SendMessage struct {
	Content string `json:"content"`
	MsgType int    `json:"msg_type"`
	MsgId   string `json:"msg_id"` // 带上该字段后发送消息变成被动消息，限制比主动消息少
}

func ExtractPrivateChatMessage(data []byte) *model.PrivateChatMessage {
	var msg model.PrivateChatMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		logging.Error("failed to unmarshal PrivateChatMessage", zap.Error(err))
		return nil
	}
	return &msg
}
func ExtractGroupMessage(data []byte) *model.GroupAtMessage {
	var msg model.GroupAtMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		logging.Error("failed to unmarshal GroupAtMessage", zap.Error(err))
		return nil
	}
	return &msg
}
