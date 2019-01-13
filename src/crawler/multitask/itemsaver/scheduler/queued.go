package scheduler

import "crawler/multitask/itemsaver/engine"

type QueuedScheduler struct {
	requestChan chan engine.Request
	// 每1个worker对外的类型是 chan engine.Request
	// 每1个worker创建1个不同的 chan engine.Request
	// 100个worker就是100个chan
	// 这100个chan如何调节 再装入1个总的chan
	workerChan chan chan engine.Request
}

// 实现 engine.Scheduler 接口的 Submit(engine.Request) 方法
func (s *QueuedScheduler)Submit(r engine.Request){
	s.requestChan <- r
}

// 实现 engine.Scheduler 接口的 ConfigureMasterWorkerChan(chan engine.Request) 方法
// func (s *QueuedScheduler)ConfigureMasterWorkerChan(in chan engine.Request){
//	panic("implement me")
// }

// 实现 engine.Scheduler 接口的 WorkerChan()chan engine.Request 方法
func (s *QueuedScheduler)WorkerChan() chan engine.Request{
	// Scheduler分配Worker 1个 chan Request
	return make(chan engine.Request)
}

// 总控调度
func (s *QueuedScheduler)Run(){
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	// 无限轮询等待新Request和Worker到来
	go func() {
		// 声明2个队列
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var acitveRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) >0 {
				activeWorker = workerQ[0]
				acitveRequest = requestQ[0]
			}
			select{
			case r := <- s.requestChan:
				// 接收到1个Request就让它排队
				requestQ = append(requestQ,r)
			case w := <- s.workerChan:
				// 接收到1个Worker也让Worker排队
				workerQ = append(workerQ,w)
			case activeWorker <- acitveRequest:
				// 从队列出拿掉Request Worker
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}

func (s *QueuedScheduler)WorkerReady(w chan engine.Request){
	// 有1个worker准备好了 可以接收Request了 worker通知Scheduler
	// Scheduler选择了这个worker对应的chan 然后进行调度给它的chan发送数据
	s.workerChan <- w
}