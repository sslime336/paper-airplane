// Package calabiYau 卡拉彼丘相关
package calabiYau

import (
	"go.uber.org/zap"
)

// 房间号是6位

var log *zap.Logger

func BindLogger(logger *zap.Logger) {
	log = logger.Named("calabiYau")
}
