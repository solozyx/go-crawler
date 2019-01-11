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