package engine

import (
	"log"
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