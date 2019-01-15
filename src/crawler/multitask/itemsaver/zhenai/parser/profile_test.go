package parser

import (
	"testing"
	"io/ioutil"
	"crawler/multitask/itemsaver/model"
	"crawler/multitask/itemsaver/engine"
)

func TestParseProfile(t *testing.T) {
	url := "http://album.zhenai.com/u/1757801149"
	body,err := ioutil.ReadFile("1757801149.html")
	if err != nil{
		panic(err)
	}
	// url := "http://album.zhenai.com/u/1757801149"
	result := ParseProfile(body,url,"等风也等你")
	actual := result.Items[0]

	expected := engine.Item{
		Url:"http://album.zhenai.com/u/1757801149",
		Type:"zhenai",
		Id:"1757801149",
		Payload:model.Profile{
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
		},
	}

	if actual != expected {
		t.Errorf("不相同的会员信息: \n profile = %v : \n right = %v", actual, expected)
	}
}

/*
func TestFetchProfile(t *testing.T) {
	if _,err := fetcher.Fetch("http://album.zhenai.com/u/1757801149"); err != nil{
		t.Error(err)
	}
}
*/