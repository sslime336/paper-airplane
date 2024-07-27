package handler

import (
	"encoding/json"
	myevnt "paper-airplane/handler/event"
	"paper-airplane/handler/group"
	prvMsg "paper-airplane/handler/privateMsg"
	"paper-airplane/logging"
	"paper-airplane/model"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"go.uber.org/zap"
)

func Get() event.PlainEventHandler {
	return func(payload *dto.WSPayload, data []byte) error {
		var g model.General
		if err := json.Unmarshal(data, &g); err != nil {
			logging.Error("failed to decode json", zap.Error(err))
			return nil
		}
		switch g.T {
		case myevnt.C2CMessageCreate:
			// 用户单聊发消息给机器人时候
			return prvMsg.Handler(payload, data)
		case myevnt.GroupAtMessageCreate:
			// 用户在群里@机器人时收到的消息
			return group.Handler(payload, data)
		}
		return nil
	}
}
