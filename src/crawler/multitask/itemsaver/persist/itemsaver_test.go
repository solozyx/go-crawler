package persist

import (
	"testing"
	"context"
	"encoding/json"
	"gopkg.in/olivere/elastic.v5"
	"crawler/multitask/itemsaver/model"
)

func TestSave(t *testing.T) {
	expected := model.Profile{
		Name:		"等风也等你",
		Gender:		"",
		Age:		25,
		Height:		162,
		Weight:		53,
		Income:		"3-5千",
		Marriage:	"未婚",
		//Education:	"大学本科",
		//Occupation:	"会计",
		Hukou:		"湖北武汉",
		//Xingzuo:	"天秤座",
		//House:		"租房",
		//Car:		"未买车",
	}
	id,err := save(expected)
	if err != nil{
		panic(err)
	}

	// TODO: Try to start up elasticsearch here using docker go client
	client,err := elastic.NewClient(
		elastic.SetURL(esAddr),
		elastic.SetSniff(false))
	if err != nil {
		panic(err)
	}
	// 获取es存储数据
	resp,err := client.Get().Index(esIndex).Type(esType).Id(id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%+v\n",resp)
	t.Logf("%s\n",resp.Source)

	var profile model.Profile
	if err = json.Unmarshal(*resp.Source,&profile); err != nil{
		panic(err)
	}
	if expected != profile {
		t.Errorf("got %v , expected %v",profile,expected)
	} else {
		t.Logf("got      %v",profile)
		t.Logf("expected %v",expected)
	}
}