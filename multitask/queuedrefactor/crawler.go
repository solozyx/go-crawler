package main

import (
	"crawler/multitask/queuedrefactor/zhenai/parser"
	"crawler/multitask/queuedrefactor/engine"
	"crawler/multitask/queuedrefactor/scheduler"
)

func main() {
	/*
	engine.SimpleEngine{}.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
	*/

	e := &engine.ConcurrentEngine{
		// Scheduler:&scheduler.SimpleScheduler{},
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
	}
	e.Run(engine.Request{
		// Url:"http://www.zhenai.com/zhenghun",
		Url:"http://www.zhenai.com/zhenghun/shanghai",
		// ParserFunc:parser.ParseCityList,
		ParserFunc:parser.ParseCity,
	})
}