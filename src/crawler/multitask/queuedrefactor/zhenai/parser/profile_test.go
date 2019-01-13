package parser

import (
	"testing"
	"io/ioutil"
	"crawler/multitask/queuedrefactor/model"
)

func TestParseProfile(t *testing.T) {
	body,err := ioutil.ReadFile("1757801149.html")
	if err != nil{
		panic(err)
	}
	// url := "http://album.zhenai.com/u/1757801149"
	result := ParseProfile(body,"等风也等你")
	profile := result.Items[0]

	right := model.Profile{
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

	if profile != right {
		t.Errorf("不相同的会员信息: \n profile = %v : \n right = %v", profile, right)
	}
}

/*
func TestFetchProfile(t *testing.T) {
	if _,err := fetcher.Fetch("http://album.zhenai.com/u/1757801149"); err != nil{
		t.Error(err)
	}
}
*/