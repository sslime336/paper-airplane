package main

import (
	"context"
	"os"
	"os/signal"
	"time"

	"github.com/sslime336/paper-airplane/bot"
	"github.com/sslime336/paper-airplane/config"
	"github.com/sslime336/paper-airplane/dao"
	"github.com/sslime336/paper-airplane/db"
	"github.com/sslime336/paper-airplane/handler"
	"github.com/sslime336/paper-airplane/keys"
	"github.com/sslime336/paper-airplane/logging"
	"github.com/sslime336/paper-airplane/service"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"go.uber.org/zap"
)

func init() {
	config.Init()
	logging.Init()
	db.Init()
	dao.SetDefault(db.Sqlite)
	handler.Init()
	service.Init()
}

func main() {
	log := logging.Named("init")

	conf := config.App.Bot
	token := token.BotToken(conf.AppId, conf.Token)
	var api openapi.OpenAPI
	if os.Getenv(keys.BotMode) == "release" {
		api = botgo.NewOpenAPI(token).WithTimeout(3 * time.Second)
	} else {
		api = botgo.NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
	}
	bot.BuildClient(api)

	botgo.SetLogger(logging.Named("bot.tencent.client").WithOptions(zap.AddCallerSkip(1)).Sugar())

	ws, err := api.WS(context.Background(), nil, "")
	if err != nil {
		log.Fatal("ws connect failed", zap.Error(err))
	}

	intent := websocket.RegisterHandlers(handler.Get())
	intent |= 1 << 25 // 注册群和私聊消息

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		log.Info("bot was exited by Ctrl C", zap.Any("signal", <-sig))
		os.Exit(0)
	}()

	if err := botgo.NewSessionManager().Start(ws, token, &intent); err != nil {
		log.Fatal("unexpected exit", zap.Error(err))
	}
}
