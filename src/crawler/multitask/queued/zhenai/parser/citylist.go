package parser

import (
	"crawler/multitask/queued/engine"
	"fmt"
	"regexp"
)

// <a href="http://www.zhenai.com/zhenghun/aba" data-v-473e2ba0="">阿坝</a>
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// contents 原网页转为UTF-8编码的text文本数据
// []Requet []Item
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	// TODO 为了便于测试 限制城市数量
//	limit := 1
	for _,m := range matches{
		result.Items = append(result.Items,"City " + string(m[2]))
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			// nil可编译运行但是调用nil()会panic ParserFunc:nil,
			// ParserFunc:engine.NilParser,
			ParserFunc:ParseCity, // 针对这里的Url使用ParseCity来解析
		})
//		limit --
//		if limit == 0 {
//			break // os.Exit(0)
//		}
	}
	fmt.Printf("Matches found citys count = %d \n",len(matches))
	return result
}