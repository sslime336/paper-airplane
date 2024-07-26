package handler

import (
	"encoding/json"
	"paper-airplane/logging"
	"paper-airplane/model"
	"strings"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"github.com/tencent-connect/botgo/openapi"
	"go.uber.org/zap"
)

var api openapi.OpenAPI

func RegisterApi(openapi openapi.OpenAPI) {
	api = openapi
}

func Get() event.PlainEventHandler {
	return func(event *dto.WSPayload, data []byte) error {
		var g model.General
		if err := json.Unmarshal(data, &g); err != nil {
			logging.Error("failed to decode json", zap.Error(err))
			return nil
		}
		switch g.T {
		// 用户单聊发消息给机器人时候
		case EventC2CMessageCreate:
			do := ExtractPrivateChatMessage(data)
			msgId := do.Data.ID
			return SendTextMessagePrivate(do.Data.Author.UserOpenid, msgId, "不是，哥们儿")
		// 用户在群里@机器人时收到的消息
		case EventGroupAtMessageCreate:
			do := ExtractGroupMessage(data)
			msgId := do.Data.ID
			if strings.HasPrefix(strings.TrimSpace(do.Data.Content), "/ping") {
				return SendTextMessageGroup(do.Data.GroupOpenid, msgId, "pong")
			}
			return SendTextMessageGroup(do.Data.GroupOpenid, msgId, "不是，哥们儿")
		}
		return nil
	}
}
