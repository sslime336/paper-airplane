package mq

import (
	"github.com/sslime336/paper-airplane/logging"
	"go.uber.org/zap"
)

var log *zap.Logger

func Init() {
	log = logging.Logger().Named("mq")

	initMemMq()
}
