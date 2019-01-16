package model

import "crawler/multitask/itemsaver/engine"

type SearchResult struct {
	// 总共找到多少个item
	Hits     int64
	// 从哪个item开始
	Start    int
	// 渲染到item的数据
	Items []engine.Item
	// Items    []interface{}

	Query    string
	PrevFrom int
	NextFrom int
}