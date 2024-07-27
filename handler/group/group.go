package group

import (
	"paper-airplane/bot"
	"paper-airplane/bot/command"

	"github.com/tencent-connect/botgo/dto"
)

func Handler(event *dto.WSPayload, data []byte) error {
	do := bot.ExtractGroupMessage(data)
	msgId := do.MsgId()
	groupOpenId := do.GroupOpenId()

	// /ping 指令
	if command.Contains(do.Content(), command.Ping) {
		return bot.PaperAirplane.ToGroup(groupOpenId).Reply(msgId, "pong")
	}

	// spark.Send(content)
	// select {
	// case ans := <-spark.Answer:
	// 	return bot.PaperAirplane.ToGroup(do.GroupOpenId()).Reply(msgId, "")
	// case <-time.After(3 * time.Second):
	// 	return bot.SendTextMessageGroup(do.Data.GroupOpenid, msgId, "timeout")
	// }
	return nil
}
