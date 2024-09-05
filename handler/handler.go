package handler

import (
	"encoding/json"

	myevnt "github.com/sslime336/paper-airplane/handler/event"
	"github.com/sslime336/paper-airplane/handler/group"
	prvMsg "github.com/sslime336/paper-airplane/handler/privateMsg"
	"github.com/sslime336/paper-airplane/logging"
	"github.com/sslime336/paper-airplane/ws"

	"github.com/tencent-connect/botgo/dto"
	"github.com/tencent-connect/botgo/event"
	"go.uber.org/zap"
)

var log *zap.Logger

func Init() {
	group.Init()
	prvMsg.Init()

	log = logging.Logger().Named("handler")
}

func Get() event.PlainEventHandler {
	return func(payload *dto.WSPayload, data []byte) error {
		var g ws.General
		if err := json.Unmarshal(data, &g); err != nil {
			log.Error("failed to decode json", zap.Error(err))
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
