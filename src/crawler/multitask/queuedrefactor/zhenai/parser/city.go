package parser

import (
	"regexp"
	"crawler/multitask/queuedrefactor/engine"
)

// <a href="http://album.zhenai.com/u/99767952" target="_blank">慕斯</a>
// const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`
var(
	profileRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`)
	// <a href="http://www.zhenai.com/zhenghun/shanghai/2">下一页</a>
	cityUrlRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)
)

func ParseCity(contents []byte) engine.ParseResult{
	// re := regexp.MustCompile(cityRe)
	matches := profileRe.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,m := range matches{
		// result.Items = append(result.Items,"User " + string(m[2]))
		name := string(m[2])
		// 用户名不生成
		// result.Items = append(result.Items,"User " + name)
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			// ParserFunc:engine.NilParser,
			ParserFunc:func(c []byte)engine.ParseResult{
							// return ParseProfile(c,string(m[2]))
							return ParseProfile(c,name)
					   },
		})
	}

	// 取本页面其它城市链接
	matches = cityUrlRe.FindAllSubmatch(contents, -1)
	for _,m := range matches{
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			ParserFunc:ParseCity,
		})
	}

	return result
}