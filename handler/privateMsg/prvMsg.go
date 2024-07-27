package prvMsg

import (
	"paper-airplane/bot"

	"github.com/tencent-connect/botgo/dto"
)

func Handler(event *dto.WSPayload, data []byte) error {
	prvMsg := bot.ExtractPrivateChatMessage(data)
	return bot.PaperAirplane.
		ToUser(prvMsg.UserOpenId()).
		Reply(prvMsg.MsgId(), "复读: "+prvMsg.Content())
}
