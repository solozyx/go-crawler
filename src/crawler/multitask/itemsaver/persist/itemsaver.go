package persist

import (
	"log"
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"github.com/pkg/errors"
	"crawler/multitask/itemsaver/engine"
)

// func ItemSaver() chan interface{} {
// index 向ES存储数据的database
func ItemSaver(esAddr,index string) (chan engine.Item,error) {
	// es client 单例
	client,err := elastic.NewClient(
		elastic.SetURL(esAddr),
		// Must turn off Sniff in docker env
		elastic.SetSniff(false))
	if err != nil{
		// panic(err)
		return nil,err
	}
	fmt.Println("ItemSaver elasticsearch connection success ")

	// ItemSaver 使用 out chan 和 Engine 通信
	// out := make(chan interface{})
	out := make(chan engine.Item)
	// real save logic
	go func() {
		itemCount := 0
		for {
			item := <- out
			itemCount ++
			log.Printf("ItemSaver : got item #%d : %v \n",itemCount,item)
			err := save(client,index,item)
			if err != nil{
				// 重试
				// 放弃 爬虫会把非常多的数据 十几万 上百万 甚至更多 其中1个item数据存错 不要紧
				log.Printf("ItemSaver error saving item %v : %v",item,err)
			}
		}
	}()
	return out,nil
}

// ElasticSearch 存储用户数据
// ES是RESTful服务可以直接用net/http包访问
// ES官方客户端
// func save(item interface{})(id string,err error){
// func save(item engine.Item) error {
func save(client *elastic.Client, index string, item engine.Item) error {
	// save data into es
	// es:9200/index(db)/type(table)/id(id)
	// Id()可以指定也可以不指定 不指定es会自动生成
	if item.Type == "" {
		return errors.New("item saver must supply type")
	}
	// resp,err := client.Index().Index(esIndex).Type(item.Type).Id(item.Id).BodyJson(item).Do(context.Background())
	indexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != ""{
		indexService.Id(item.Id)
	}
	resp,err := indexService.Do(context.Background())
	if err != nil{
		return nil
	}
	fmt.Printf("%+v\n",resp)
	return nil
}