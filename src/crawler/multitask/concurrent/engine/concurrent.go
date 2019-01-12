package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	// interface{}定义方法不需要参数名 指明参数类型即可
	Submit(Request)
	// 把Run方法生成的 in := make(chan Request) 放到接口中
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine)Run(seeds ...Request){
	// 所有worker共用1个input/output channel
	in := make(chan Request)
	e.Scheduler.ConfigureMasterWorkerChan(in)
	out := make(chan ParseResult)

	for i := 0; i < e.WorkerCount; i++{
		createWorker(in,out)
	}

	for _,r := range seeds{
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	// 程序没有退出条件 一直轮询等待新数据
	for {
		result := <- out
		for _,item := range result.Items{
			itemCount ++
			log.Printf("Got #%d : item = %v \n",itemCount,item)
		}
		for _,request := range result.Requests{
			// 这里结构体值传递
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult){
	go func() {
		for{
			request := <- in
			result,err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}