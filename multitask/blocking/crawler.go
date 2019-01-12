package main

import (
	"crawler/multitask/blocking/zhenai/parser"
	"crawler/multitask/blocking/scheduler"
	"crawler/multitask/blocking/engine"
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
		WorkerCount:3,
	}
	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
}