package main

import (
	"net/http"
	"crawler/multitask/itemsaver/frontend/controller"
)

func main() {
	webDir := "./view"
	tpl := "./view/template.html"
	// 静态资源展示 非 /search 就展示静态资源
	http.Handle("/",http.FileServer(http.Dir(webDir)))
	// http.Handle("/search",controller.SearchResultHandler{})
	http.Handle("/search",controller.CreateSearchResultHandler(tpl))
	err := http.ListenAndServe(":8888",nil)
	if err != nil{
		panic(err)
	}
}