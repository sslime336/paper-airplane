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
	"github.com/sslime336/paper-airplane/logging"
	"github.com/sslime336/paper-airplane/service"
	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/openapi"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"go.uber.org/zap"
)

var log *zap.Logger

func main() {
	var devMode bool
	if os.Getenv("BOT_MODE") == "release" {
		devMode = true
	}
	conf := config.ParseConfig[config.App]("./config.yaml")
	logging.Init(conf.Log.Path, "bot.log", devMode)
	db.Init(&conf)
	dao.SetDefault(db.Sqlite)
	handler.Init()
	service.Init(&conf)
	log = logging.Named("bot.init")

	token := token.BotToken(conf.Bot.AppId, conf.Bot.Token)
	log.Info("bot start", zap.String("token", token.GetString()))

	var api openapi.OpenAPI
	if devMode {
		api = botgo.NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
	} else {
		api = botgo.NewOpenAPI(token).WithTimeout(3 * time.Second)
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
