package engine

import (
	"log"
	"crawler/singletask/fetcher"
)

func Run(seeds ...Request) {
	// Engine 维护1个Request队列
	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s",r.Url)

		// Fetch每1个Request获得原始网页转换为UTF-8编码的原始文本数据
		body,err := fetcher.Fetch(r.Url)
		if err != nil{
			log.Printf("Fetcher error fetching url = %s,err = %v",r.Url,err)
			continue
		}

		// 解析原始网页文本数据
		parseResult := r.ParserFunc(body)
		// 添加parseResult所有的Request到requests
		requests = append(requests,parseResult.Requests...)
		for _,item := range parseResult.Items {
			log.Printf("Got item = %v",item)
		}
	}

	log.Printf("Engine exit :( ")
}