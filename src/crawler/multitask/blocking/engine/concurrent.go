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
	// interface method
	e.Scheduler.ConfigureMasterWorkerChan(in)
	out := make(chan ParseResult)

	for i := 0; i < e.WorkerCount; i++{
		log.Printf("Engine 在 main协程 开启 WorkerCount = %d 个子协程, goroutine No = %d",e.WorkerCount,i)
		createWorker(in,out)
	}

	for _,r := range seeds{
		log.Printf("Engine 在 main协程 向 in chan 投递 Request 解除 Worker 子协程阻塞")
		// interface use
		e.Scheduler.Submit(r)
	}

	for {
		result := <- out
		log.Printf("Engine main协程 消费 out chan 解除 Worker 子协程阻塞")
		for _,item := range result.Items{
			log.Printf("Got item = %v \n",item)
		}
		for _,request := range result.Requests{
			// 这里结构体值传递
			log.Printf("Engine main协程 阻塞 等待Worker子协程消费 in chan ...")
			// interface use
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request,out chan ParseResult){
	go func() {
		for{
			log.Printf("Worker 子协程阻塞...")
			request := <- in
			log.Printf("Worker 子协程解除阻塞,开始爬取网页")
			result,err := worker(request)
			if err != nil{
				continue
			}
			log.Printf("Worker 子协程 爬取网页结束 返回main协程 result")
			log.Printf("Worker 子协程 阻塞 等待 main协程 消费 out chan ...")
			out <- result

			//select {
			//case out<-result:
			//default:
			//	log.Printf("Worker blocking...\n")
			//}
		}
	}()
}