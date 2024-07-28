package service

import "paper-airplane/service/spark"

type SparkAI struct {
	*spark.Session
}

func Init() {
	spark.Init()
}
