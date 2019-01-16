package controller

import (
	"net/http"
	"strings"
	"strconv"
	"fmt"
	"context"
	"reflect"
	"regexp"
	"gopkg.in/olivere/elastic.v5"
	"crawler/multitask/itemsaver/frontend/view"
	"crawler/multitask/itemsaver/conf"
	"crawler/multitask/itemsaver/frontend/model"
	"crawler/multitask/itemsaver/engine"
)

type SearchResultHandler struct {
	view view.SearchResultView
	client *elastic.Client
}

// 静态模板文件名
func CreateSearchResultHandler(tpl string) SearchResultHandler {
	client,err := elastic.NewClient(elastic.SetURL(conf.EsAddr),elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}
	return SearchResultHandler{
		view:view.CreateSearchResultView(tpl),
		client:client,
	}
}

// localhost:8888/search?q=男 已购房&from=20
func (h SearchResultHandler)ServeHTTP(w http.ResponseWriter,req *http.Request){
	q := strings.TrimSpace(req.FormValue("q"))

	from,err := strconv.Atoi(req.FormValue("from"))
	if err != nil{
		from = 0
	}
	// fmt.Fprintf(w,"q= %s, from= %d",q,from)

	var page model.SearchResult
	page,err = h.getSearchResult(q,from)
	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}

	err = h.view.Render(w,page)
	if err != nil{
		http.Error(w,err.Error(),http.StatusBadRequest)
	}
}

func (h SearchResultHandler)getSearchResult(q string,from int)(model.SearchResult,error){
	var searchResult model.SearchResult
	searchResult.Query = q

	resp,err := h.client.
		Search(conf.EsIndex).
		Query(elastic.NewQueryStringQuery(rewriteQueryString(q))).
		From(from).
		Do(context.Background())

	if err != nil{
		return searchResult,err
	}

	searchResult.Hits = resp.TotalHits()
	searchResult.Start = from
	for _,v := range resp.Each(reflect.TypeOf(engine.Item{})){
		item := v.(engine.Item)
		searchResult.Items = append(searchResult.Items,item)
	}
	searchResult.PrevFrom = searchResult.Start - len(searchResult.Items)
	searchResult.NextFrom = searchResult.Start + len(searchResult.Items)
	fmt.Printf("%+v",searchResult)
	return searchResult,nil
}

func rewriteQueryString(q string) string {
	// Payload.Age:(<30)
	re := regexp.MustCompile(`([A-Z][a-z]*):`)
	// 正则表达式 可以 查找 和 replace
	// 字段名 $1
	return  re.ReplaceAllString(q,"Payload.$1:")
}

