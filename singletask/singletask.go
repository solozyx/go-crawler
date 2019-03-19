package main

import (
	"crawler/singletask/engine"
	"crawler/singletask/zhenai/parser"
)

func main() {
	seedReq := engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		ParserFunc:parser.ParseCityList,
	}
	engine.Run(seedReq)
}