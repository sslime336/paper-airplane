package main

import (
	"context"
	"paper-airplane/config"
	"paper-airplane/handler"
	"paper-airplane/logging"
	"time"

	"github.com/tencent-connect/botgo"
	"github.com/tencent-connect/botgo/token"
	"github.com/tencent-connect/botgo/websocket"
	"go.uber.org/zap"
)

func init() {
	config.Init()
	logging.Init()
}

func main() {
	conf := config.Bot()
	token := token.BotToken(conf.AppId, conf.Token)
	api := botgo.NewSandboxOpenAPI(token).WithTimeout(3 * time.Second)
	handler.RegisterApi(api)

	botgo.SetLogger(logging.Logger().Sugar())

	ws, err := api.WS(context.Background(), nil, "")
	if err != nil {
		logging.Fatalf("%+v, error:%v", ws, err)
	}

	intent := websocket.RegisterHandlers(handler.Get())
	intent |= 1 << 25 // 注册群和私聊消息
	if err := botgo.NewSessionManager().Start(ws, token, &intent); err != nil {
		logging.Fatal("unexpected exit", zap.Error(err))
	}
}
