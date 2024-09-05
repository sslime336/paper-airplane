// Package spark 接入了讯飞模型，主要是 SparkLite 以提供 AI 聊天对话支持
package spark

import (
	"github.com/sslime336/paper-airplane/bot"
	"github.com/sslime336/paper-airplane/config"
	"github.com/sslime336/paper-airplane/service/spark/req"
	"go.uber.org/zap"
)

var log *zap.Logger
var authUrl func() string

const HostUrlSparkLite = "wss://spark-api.xf-yun.com/v1.1/chat"

func Init(conf *config.App, logger *zap.Logger) {
	initSessionCache()
	log = logger.Named("spark")
	authUrl = func() string {
		return buildAuthUrl(HostUrlSparkLite, conf.Spark.ApiKey, conf.Spark.ApiSecret)
	}
	createSparkLiteRequest = func() *req.Request {
		return req.NewSparkLiteRequest(conf.Spark.AppId)
	}
}

func Chat(openId, msgId, msg string) error {
	log.Debug("received chat message", zap.String("message", msg))
	sess, err := NewSparkSession(openId)
	if err != nil {
		return err
	}

	if err := sess.Send(msg); err != nil {
		log.Error("send message failed", zap.Error(err))
		return bot.MyBot.ToGroup(openId).Reply(msgId, err.Error())
	}
	res, err := sess.Read()
	if err != nil {
		log.Error("read from spark failed", zap.Error(err))
		return bot.MyBot.ToGroup(openId).Reply(msgId, err.Error())
	}
	return bot.MyBot.ToGroup(openId).Reply(msgId, res)
}
