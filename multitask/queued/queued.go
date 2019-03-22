package main

import (
	"net/http"
	_ "net/http/pprof"

	"crawler/multitask/queued/engine"
	"crawler/multitask/queued/scheduler"
	"crawler/multitask/queued/zhenai/parser"
)

func main() {
	/*
	engine.SimpleEngine{}.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
	*/

	// pprof
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	// start crawler
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