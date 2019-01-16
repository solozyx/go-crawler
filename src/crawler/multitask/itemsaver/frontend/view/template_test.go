package view

import (
	"testing"
	"html/template"
	"os"
	"crawler/multitask/itemsaver/frontend/model"
	"crawler/multitask/itemsaver/engine"
	profile "crawler/multitask/itemsaver/model"
	"crawler/multitask/itemsaver/conf"
)

func TestTemplate(t *testing.T){
	out,err := os.Create("template.test.html")
	if err != nil {
		panic(err)
	}

	tpl := template.Must(template.ParseFiles("template.html"))
	page := model.SearchResult{}
	page.Hits = 123

	item := engine.Item{
		Url:"http://album.zhenai.com/u/1757801149",
		Type:conf.EsType,
		Id:"1757801149",
		Payload:profile.Profile{
			Name:     "等风也等你",
			Gender:   "",
			Age:      25,
			Height:   162,
			Weight:   53,
			Income:   "3-5千",
			Marriage: "未婚",
			//Education:	"大学本科",
			//Occupation:	"会计",
			Hukou: "湖北武汉",
			//Xingzuo:	"天秤座",
			//House:		"租房",
			//Car:		"未买车",
		},
	}
	for i := 0; i < 10; i++ {
		page.Items = append(page.Items,item)
	}

	// err := tpl.Execute(os.Stdout,page)
	err = tpl.Execute(out,page)
	if err != nil{
		panic(err)
	}
}