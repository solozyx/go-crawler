package engine

type Request struct {
	Url string
	ParserFunc func([]byte)ParseResult
	// ParserFunc func([]byte,map[string]string)ParseResult
}

type ParseResult struct {
	Requests []Request
	// interface{}是任何类型
	Items []interface{}
}

func NilParser([]byte)ParseResult{
	return ParseResult{}
}