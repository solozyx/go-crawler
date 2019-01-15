package model

import "encoding/json"

// 爬虫项目不局限于某个特定网站 可以爬取各种网站
// 每个网站的爬取作为1个业务 最后爬取到的数据 希望能够汇总起来
// 汇总到统一的数据库 让数据更有价值

// 独立结构 与某一特定网站没有特别强的相关性
type Profile struct {
	// 昵称
	Name string
	// 性别
	Gender string
	// 年龄
	Age int
	// 身高
	Height int
	// 体重
	Weight int
	// 收入 5000-8000
	Income string
	// 婚姻状况
	Marriage string
	// 教育
	Education string
	// 职业
	Occupation string
	// 户口
	Hukou string
	// 星座
	Xingzuo string
	// 房产
	House string
	// 汽车
	Car string
}

func FromJsonObj(o interface{}) (Profile,error) {
	var p Profile
	strData,err := json.Marshal(o)
	if err != nil{
		return p,err
	}
	err = json.Unmarshal(strData,&p)
	return p,nil

}