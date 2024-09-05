package service

import (
	"github.com/sslime336/paper-airplane/config"
	"github.com/sslime336/paper-airplane/logging"
	"github.com/sslime336/paper-airplane/service/calabiYau"
	"github.com/sslime336/paper-airplane/service/spark"
	"go.uber.org/zap"
)

var log *zap.Logger

func Init(conf *config.App) {
	log = logging.Named("service")

	spark.Init(conf, log)

	calabiYau.BindLogger(log)
}
