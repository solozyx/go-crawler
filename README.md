# go-crawler
go distributed crawler using net/http regexp elastic-search

## 单任务版爬虫架构
![image](https://github.com/solozyx/go-crawler/blob/master/screenshots/singletask.png)

### 单任务版爬虫架构说明
* 给引擎Engine传入待抓取网站初始化种子seed
* 对每一个待抓取URL封装 Request{Url:string,ParserFunc func([]byte)ParseResult} 任务,Engine维护Requests任务队列
* Engine调度Requests任务队列,Fetcher(Url)抓取网页文本,转为UTF-8编码,返回给Engine
* Engine把Fetcher返回的UTF-8编码的网页原始文本数据交给Parser解析,Parser解析后返回Engine解析结果 ParseResult{Requests []Request,Items []interface{}}
* Engine把Parser返回的[]Request添加到任务队列,打印Items数据存储展示
* 网络使用率每秒 70-80K 爬取效率低下

## 并发版爬虫架构
### 简单调度器
![image](https://github.com/solozyx/go-crawler/blob/master/screenshots/multitask.png)
* GOPATH/multitask/blocking    演示了blocking阻塞

### 并发分发调度器
![image](https://github.com/solozyx/go-crawler/blob/master/screenshots/concurrentscheduler.png)
* GOPATH/multitask/concurrent  使用100个其他子协程解除blocking阻塞 使用100个Worker子协程
```
func (s *SimpleScheduler)Submit(r engine.Request){
	go func(){
		s.workerChan <- r
	}()
}
```
为每个Request创建1个goroutine,这里每个goroutine只做1件事情往Worker子goroutine使用的in chan Request 投递 Request 解除阻塞,每个分发Request的goroutine做完分发任务就退出结束.

但是流量过大,导致爬虫引擎被对方网站禁掉,引入简单的流控机制,增加rateLimiter在流量不被目标网站禁掉的情况下 网络使用率达到每秒 1.5M-1.6M
```
package fetcher
// 10毫秒执行一次请求
var rateLimiter = time.Tick(10 * time.Millisecond)
func Fetch(url string) ([]byte, error) {
	<- rateLimiter
	//...
}
```

### 队列调度器
![image](https://github.com/solozyx/go-crawler/blob/master/screenshots/queued.png)
* 并发分发调度器对项目的控制力度比较小，启动了200个goroutine，100个Worker goroutine，每个Request对应创建1条goroutine分发
分发出去的Request就收不回来了，也不知道分发出去的Request在外面怎么样了，所有的Worker都在抢同1个 in chan Request 过来的 Request 也没有办法去控制想给到哪个Worker，做一些负载均衡之类的事情无法做
1. 把Request放到Request队列，分发队列头的Request
2. Request分发给Worker，加大对Worker的控制，可以自己选择Worker做任务，引入Worker队列
3. 有了Request 和 Worker ，把自己选择的Request发给自己选择的Worker，控制力度增大
4. 100个Worker子goroutine爬取网页文本数据正则计算,1个子goroutine调度队列,大幅减少子goroutine数量
5. 去掉流控限制最大流量峰值也有 4M-5M 速度非常快,也被目标网站禁掉了, 引入队列的实现和并发分发实现性能差不多,控制力度提高
