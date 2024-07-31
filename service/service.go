package service

import (
	"github.com/sslime336/paper-airplane/service/calabiYau"
	"github.com/sslime336/paper-airplane/service/spark"
)

func Init() {
	spark.Init()
	calabiYau.Init()
}
