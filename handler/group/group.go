package group

import (
	"paper-airplane/bot"
	"paper-airplane/bot/command"
	"paper-airplane/logging"
	"paper-airplane/service/spark"

	"github.com/tencent-connect/botgo/dto"
	"go.uber.org/zap"
)

var sessionMap = map[string]*spark.Session{}

func Handler(event *dto.WSPayload, data []byte) error {
	do := bot.ExtractGroupMessage(data)
	msgId := do.MsgId()
	groupOpenId := do.GroupOpenId()
	memberOpenId := do.D.Author.MemberOpenid

	// `/ping` 指令
	if command.Contains(do.RawContent(), command.Ping) {
		return bot.PaperAirplane.ToGroup(groupOpenId).Reply(msgId, "pong")
	}

	var session *spark.Session
	if sess, ok := sessionMap[memberOpenId]; ok {
		session = sess
	} else {
		sessionMap[memberOpenId], _ = spark.NewSparkSession()
		session = sessionMap[memberOpenId]
	}

	if command.Contains(do.RawContent(), command.Chat) {
		session.Send(do.Content())
		res, err := session.Read()
		if err != nil {
			logging.Error("read from spark failed", zap.Error(err))
			return bot.PaperAirplane.ToGroup(groupOpenId).Reply(msgId, "Spark 读取错误")
		}
		return bot.PaperAirplane.ToGroup(groupOpenId).Reply(msgId, res)
	}

	return nil
}
