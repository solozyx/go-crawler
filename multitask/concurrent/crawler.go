package main

import (
	"crawler/multitask/concurrent/zhenai/parser"
	"crawler/multitask/concurrent/scheduler"
	"crawler/multitask/concurrent/engine"
)

func main() {
	/*
	engine.SimpleEngine{}.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
	*/

	e := &engine.ConcurrentEngine{
		Scheduler:&scheduler.SimpleScheduler{},
		WorkerCount:100,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}