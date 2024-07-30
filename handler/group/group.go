package group

import (
	"paper-airplane/bot"
	"paper-airplane/bot/command"
	"paper-airplane/logging"
	"paper-airplane/service/spark"

	"github.com/tencent-connect/botgo/dto"
	"go.uber.org/zap"
)

var log *zap.Logger

func Init() {
	log = logging.Logger().Named("handler.group")
}

func Handler(event *dto.WSPayload, data []byte) error {
	do := bot.ExtractGroupMessage(data)
	log.Debug("received atGroupMessage", zap.Any("model.GroupAtMessage", *do))

	msgId := do.MsgId()
	groupOpenId := do.GroupOpenId()

	if cmd, ok := bot.PaperAirplane.ParseCommand(do.RawContent()); ok {
		switch cmd {
		case command.Ping:
			return bot.PaperAirplane.ToGroup(groupOpenId).Reply(msgId, "pong")
		}
	}
	return spark.Chat(groupOpenId, msgId, do.RawContent())
}
