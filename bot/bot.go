package bot

import (
	"os"
	"strings"

	"github.com/sslime336/paper-airplane/bot/botcmd"
	"github.com/sslime336/paper-airplane/keys"
	"github.com/sslime336/paper-airplane/logging"
	"github.com/tencent-connect/botgo/openapi"
	"go.uber.org/zap"
)

type Bot struct {
	openapi.OpenAPI
	hostUrl string
}

var hostUrlTemplate struct {
	User  string
	Group string
}

var MyBot *Bot

var log *zap.Logger

const (
	hostUser         = "https://api.sgroup.qq.com/v2/users/%s/messages"
	hostGroup        = "https://api.sgroup.qq.com/v2/groups/%s/messages"
	hostUserSandbox  = "https://sandbox.api.sgroup.qq.com/v2/users/%s/messages"
	hostGroupSandbox = "https://sandbox.api.sgroup.qq.com/v2/groups/%s/messages"
)

func BuildClient(api openapi.OpenAPI) {
	MyBot = new(Bot)
	MyBot.OpenAPI = api

	hostUrlTemplate.User = hostUserSandbox
	hostUrlTemplate.Group = hostGroupSandbox

	if os.Getenv(keys.BotMode) == "release" {
		hostUrlTemplate.User = hostUser
		hostUrlTemplate.Group = hostGroup
	}

	log = logging.Named("bot")
}

func (b *Bot) ParseCommand(content string) (botcmd.Command, bool) {
	ctnt := strings.TrimSpace(content)
	log.Debug("bot received message", zap.String("trimed-content", content))
	fields := strings.Split(ctnt, " ")
	log.Debug("bot received message fields", zap.Strings("message-fields", fields))

	cmd, ok := botcmd.CommandMap[fields[0]]
	return cmd, ok
}
