package service

import "paper-airplane/service/spark"

var Spark struct {
	Send func(string) error
}

func Init() {
	spark.Init()
	Spark.Send = spark.Send
}
