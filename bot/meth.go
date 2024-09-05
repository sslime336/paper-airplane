package bot

import (
	"context"
	"fmt"
)

func (bot *Bot) Message(msg Message) error {
	_, err := bot.Transport(context.Background(), "POST", bot.hostUrl, msg)
	return err
}

// Send 发送文本消息
func (bot *Bot) Send(msg string) error {
	return bot.Message(TextMessage(msg, ""))
}

// Reply 发送被动文本消息
func (bot *Bot) Reply(msgId, msg string) error {
	return bot.Message(TextMessage(msg, msgId))
}

func (bot *Bot) ToUser(userOpenId string) *Bot {
	bot.hostUrl = fmt.Sprintf(hostUrlTemplate.User, userOpenId)
	return bot
}

func (bot *Bot) ToGroup(groupOpenId string) *Bot {
	bot.hostUrl = fmt.Sprintf(hostUrlTemplate.Group, groupOpenId)
	return bot
}
