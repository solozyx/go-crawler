package main

import (
	"crawler/multitask/queued/zhenai/parser"
	"crawler/multitask/queued/scheduler"
	"crawler/multitask/queued/engine"
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
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}