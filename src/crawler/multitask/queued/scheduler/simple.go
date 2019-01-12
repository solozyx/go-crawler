package scheduler

import "crawler/multitask/queued/engine"

type SimpleScheduler struct{
	workerChan chan engine.Request
}

// 实现 engine.Scheduler 接口的 Submit(engine.Request) 方法
func (s *SimpleScheduler)Submit(r engine.Request){
	// send Request down to Worker chan
	// 该行代码会在main协程阻塞
	// s.workerChan <- r

	// goroutine 解除阻塞
	go func(){
		s.workerChan <- r
	}()
}

// 实现 engine.Scheduler 接口的 ConfigureMasterWorkerChan(chan engine.Request) 方法
// 拿到 engine.Run() 方法创建的 in := make(chan Request)
func (s *SimpleScheduler)ConfigureMasterWorkerChan(in chan engine.Request){
	s.workerChan = in
}