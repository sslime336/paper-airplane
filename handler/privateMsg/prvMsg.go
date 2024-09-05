package prvMsg

import (
	"github.com/sslime336/paper-airplane/bot"
	"github.com/sslime336/paper-airplane/logging"
	"github.com/tencent-connect/botgo/dto"
	"go.uber.org/zap"
)

var log *zap.Logger

func Init() {
	log = logging.Logger().Named("handler.group")
}

func Handler(event *dto.WSPayload, data []byte) error {
	prvMsg := bot.ExtractPrivateChatMessage(data)
	return bot.MyBot.
		ToUser(prvMsg.UserOpenId()).
		Reply(prvMsg.MsgId(), "复读: "+prvMsg.Content())
}
