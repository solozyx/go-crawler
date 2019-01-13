package engine

import (
	"log"
	"crawler/multitask/queuedrefactor/fetcher"
)

type SimpleEngine struct{

}

func (e SimpleEngine)Run(seeds ...Request) {
	// Engine 维护1个Request队列
	var requests []Request
	for _,r := range seeds{
		requests = append(requests,r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		parseResult,err := worker(r)
		if err != nil{
			continue
		}
		// 添加parseResult所有的Request到requests
		requests = append(requests,parseResult.Requests...)
		for _,item := range parseResult.Items {
			log.Printf("Got item = %v",item)
		}
	}

	log.Printf("Engine exit :( ")
}

// func (SimpleEngine)worker(r Request) (ParseResult,error) {
func worker(r Request) (ParseResult,error) {
	log.Printf("Fetching %s",r.Url)
	// Fetch每1个Request获得原始网页转换为UTF-8编码的原始文本数据
	body,err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("Fetcher error fetching url = %s,err = %v",r.Url,err)
		return ParseResult{},err
	}
	// 解析原始网页文本数据
	return r.ParserFunc(body),nil
}