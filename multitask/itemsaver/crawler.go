package main

import (
	"crawler/multitask/itemsaver/zhenai/parser"
	"crawler/multitask/itemsaver/engine"
	"crawler/multitask/itemsaver/scheduler"
	"crawler/multitask/itemsaver/persist"
	"crawler/multitask/itemsaver/conf"
)

func main() {
	/*
	engine.SimpleEngine{}.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	})
	*/

	itemChan,err := persist.ItemSaver(conf.EsAddr,conf.EsIndex)
	if err != nil {
		// client connect to es err
		panic(err)
	}

	e := &engine.ConcurrentEngine{
		// Scheduler:&scheduler.SimpleScheduler{},
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		// ItemSaver 使用 out chan 和 Engine 通信 Engine在这里做配置
		// ItemChan:persist.ItemSaver(),
		ItemChan:itemChan,
	}

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		//Url:"http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc:parser.ParseCityList,
		//ParserFunc:parser.ParseCity,
	})
}