package persist

import (
	"log"
	"context"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
)

const(
	esAddr = "http://192.168.234.142:9200"
	esIndex = "dating_profile"
	esType = "zhenai"
)

func ItemSaver() chan interface{} {
	// ItemSaver 使用 out chan 和 Engine 通信
	out := make(chan interface{})
	// real save logic
	go func() {
		itemCount := 0
		for {
			item := <- out
			itemCount ++
			log.Printf("ItemSaver : got item #%d : %v \n",itemCount,item)
			save(item)
		}
	}()
	return out
}

// ElasticSearch 存储用户数据
// ES是RESTful服务可以直接用net/http包访问
// ES官方客户端
func save(item interface{})(id string,err error){
	client,err := elastic.NewClient(
		elastic.SetURL(esAddr),
		// Must turn off Sniff in docker env
		elastic.SetSniff(false))
	if err != nil{
		// panic(err)
		return "",nil
	}
	fmt.Println("ItemSaver elasticsearch connection success ")

	// save data into es
	// es:9200/index(db)/type(table)/id(id)
	// Id()可以指定也可以不指定 不指定es会自动生成
	resp,err := client.Index().
		Index(esIndex).
		Type(esType). // Id().
		BodyJson(item).
		Do(context.Background())
	if err != nil{
		return "",nil
	}
	fmt.Printf("%+v\n",resp)
	return resp.Id,nil
}