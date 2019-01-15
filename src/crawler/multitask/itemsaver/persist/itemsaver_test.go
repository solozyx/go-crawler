package persist

import (
	"testing"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/olivere/elastic.v5"
	"crawler/multitask/itemsaver/model"
	"crawler/multitask/itemsaver/engine"
	"crawler/multitask/itemsaver/conf"
)



func TestSave(t *testing.T) {
	expected := engine.Item{
		Url:"http://album.zhenai.com/u/1757801149",
		Type:conf.EsType,
		Id:"1757801149",
		Payload:model.Profile{
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

	// TODO: Try to start up elasticsearch here using docker go client
	// es client 单例
	client,err := elastic.NewClient(
		elastic.SetURL(conf.EsAddr),
		// Must turn off Sniff in docker env
		elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}
	fmt.Println("ItemSaver elasticsearch connection success ")

	// save expected
	err = save(client,conf.EsIndex,expected)
	if err != nil{
		panic(err)
	}

	// fetch expected 获取es存储数据
	resp,err := client.Get().Index(conf.EsIndex).Type(conf.EsType).Id(expected.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%+v\n",resp)
	t.Logf("%s\n",resp.Source)

	var profile engine.Item
	if err = json.Unmarshal(*resp.Source,&profile); err != nil{
		panic(err)
	}
	var profilePayload model.Profile
	profilePayload,_ = model.FromJsonObj(profile.Payload)
	profile.Payload = profilePayload

	// verify
	if expected != profile {
		t.Errorf("got %v , expected %v",profile,expected)
	} else {
		t.Logf("got      %+v",profile)
		t.Logf("expected %+v",expected)
	}
}