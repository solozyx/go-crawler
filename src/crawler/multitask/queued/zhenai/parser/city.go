package parser

import (
	"regexp"
	"crawler/multitask/queued/engine"
)

// <a href="http://album.zhenai.com/u/99767952" target="_blank">慕斯</a>
const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)" [^>]*>([^<]+)</a>`

func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents,-1)
	result := engine.ParseResult{}
	for _,m := range matches{
		// result.Items = append(result.Items,"User " + string(m[2]))
		name := string(m[2])
		result.Items = append(result.Items,"User " + name)
		result.Requests = append(result.Requests,engine.Request{
			Url:string(m[1]),
			// ParserFunc:engine.NilParser,
			ParserFunc:func(c []byte)engine.ParseResult{
							// return ParseProfile(c,string(m[2]))
							return ParseProfile(c,name)
					   },
		})
	}
	return result
}