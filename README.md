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

### 待存储数据并发分发
![image](https://github.com/solozyx/go-crawler/blob/master/screenshots/itemsaver.png)
* Itemsaver存储数据的速度 远远快于 Worker拿到Request爬取网页获取数据的速度
从目标网站获取网络数据比本地的数据存储要慢很多,从网络拉数据,去重,rateLimiter限流不能拿太快
存储是自己能控制的,开足马力去存储
因此假设，存储Item的速度比从网上获取用户数据的速度快很多,为每个Item创建goroutine之后,Item会很快被消费掉,使用`SimpleScheduler`调度器即可
为每个Item开1个goroutine `消费Item的速度 >> 生成Item的速度` 所以开出来的goroutine不会太多
极端情况,爬虫运行项目挺久了1-2分钟也就拿到10万个人的数据,最多也就开 10万个goroutine而已,goroutine开 10万个 而且这么简单的goroutine 在性能上的代价是非常小的
```
go func(){ItemChan <- Item}()
```

