package bot

import (
	"context"
	"fmt"
)

func (bot *PaperAirplaneBot) Message(msg Message) error {
	_, err := bot.Transport(context.Background(), "POST", bot.hostUrl, msg)
	return err
}

// Send 发送文本消息
func (bot *PaperAirplaneBot) Send(msg string) error {
	return bot.Message(TextMessage(msg, ""))
}

// Reply 发送被动文本消息
func (bot *PaperAirplaneBot) Reply(msgId, msg string) error {
	return bot.Message(TextMessage(msg, msgId))
}

func (bot *PaperAirplaneBot) ToUser(userOpenId string) *PaperAirplaneBot {
	bot.hostUrl = fmt.Sprintf(hostUrlTemplate.User, userOpenId)
	return bot
}

func (bot *PaperAirplaneBot) ToGroup(groupOpenId string) *PaperAirplaneBot {
	bot.hostUrl = fmt.Sprintf(hostUrlTemplate.Group, groupOpenId)
	return bot
}
