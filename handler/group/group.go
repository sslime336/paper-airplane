package group

import (
	"github.com/sslime336/paper-airplane/bot"
	"github.com/sslime336/paper-airplane/bot/botcmd"
	"github.com/sslime336/paper-airplane/logging"
	"github.com/sslime336/paper-airplane/service/spark"
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

	if cmd, ok := bot.MyBot.ParseCommand(do.RawContent()); ok {
		switch cmd {
		case botcmd.Ping:
			return bot.MyBot.ToGroup(groupOpenId).Reply(msgId, "pong")
		}
	}
	return spark.Chat(groupOpenId, msgId, do.RawContent())
}
