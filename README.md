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
