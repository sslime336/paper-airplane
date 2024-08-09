package service

import (
	"github.com/sslime336/paper-airplane/logging"
	"github.com/sslime336/paper-airplane/service/calabiYau"
	"github.com/sslime336/paper-airplane/service/spark"
	"go.uber.org/zap"
)

var log *zap.Logger

func Init() {
	log = logging.Named("service")

	spark.BindLogger(log)
	spark.Init()

	calabiYau.BindLogger(log)
}
