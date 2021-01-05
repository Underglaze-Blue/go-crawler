package main

import (
	"crawler/engine"
	"crawler/parsist"
	"crawler/scheduler"
	"crawler/zhenai/parser"
)

func main() {
	saver, err := parsist.ItemSaver("dating_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 100,
		ItemChan:    saver,
	}
	e.Run(engine.Request{
		Url:        "http://localhost:8080/mock/www.zhenai.com/zhenghun/baotou",
		ParserFunc: parser.ParseCity,
	})
}
