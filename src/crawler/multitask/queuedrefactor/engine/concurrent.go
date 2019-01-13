package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
}

type Scheduler interface {
	// Scheduler组合 ReadyNotifer 接口
	ReadyNotifer

	// interface{}定义方法不需要参数名 指明参数类型即可
	// 分发Engine所在协程创建的Request任务
	Submit(Request)

	// 把Engine.Run方法生成的 in := make(chan Request) 放到接口中
	// refactor[:1]
	// ConfigureMasterWorkerChan(chan Request)

	// Worker可进行下一轮读取chan操作就绪状态通知
	// refactor[:2] Scheduler接口4个方法比较重
	// 将该接口方法拿到 ReadyNotifer 接口
	// Scheduler组合 ReadyNotifer 接口
	// WorkerReady(chan Request)

	// 总控调度
	Run()

	// refactor[:1]
	// Engine问Scheduler 我有1个worker 调度给我哪个 chan Request
	WorkerChan() chan Request
}

type ReadyNotifer interface{
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine)Run(seeds ...Request){
	// 所有worker共用1个input/output channel
	// in := make(chan Request)
	// e.Scheduler.ConfigureMasterWorkerChan(in)

	out := make(chan ParseResult)
	// 执行调度器
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++{
		// createWorker(in,out)
		// createWorker(out,e.Scheduler)

		// refactor[:1]
		// Worker向Scheduler要 chan Request 来做任务
		createWorker(e.Scheduler.WorkerChan(),out,e.Scheduler)
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

// func createWorker(in chan Request,out chan ParseResult){
// func createWorker(out chan ParseResult,s Scheduler){
// func createWorker(in chan Request,out chan ParseResult,s Scheduler){
func createWorker(in chan Request,out chan ParseResult,ready ReadyNotifer){
	// in := make(chan Request)
	go func() {
		for{
			// tell Scheduler i'm ready
			// 每个Worker有1个chan Request
			// 通知Scheduler本Worker就绪
			// s.WorkerReady(in)
			ready.WorkerReady(in)
			// Scheduler给本Worker投递分发1个Request 触发Worker工作
			request := <- in
			result,err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}