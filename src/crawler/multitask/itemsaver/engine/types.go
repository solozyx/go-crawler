package engine

type Request struct {
	Url string
	ParserFunc func([]byte)ParseResult
	// ParserFunc func([]byte,map[string]string)ParseResult
}

type Item struct{
	// 通用参数
	// 用户profile url 商品detail url 文章detail url ...
	Url string
	// 用户 商品 文章 ... id 在ES存储时用于去重
	Id string
	// Type (table) 表示爬取的哪个目标网站
	Type string

	// 具体业务参数
	Payload interface{}
}

type ParseResult struct {
	Requests []Request
	// interface{}是任何类型
	Items []Item //[]interface{}
}

func NilParser([]byte)ParseResult{
	return ParseResult{}
}