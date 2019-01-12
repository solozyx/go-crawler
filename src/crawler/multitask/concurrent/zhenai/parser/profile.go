package parser

import (
	"regexp"
	"strconv"
	"crawler/multitask/concurrent/engine"
	"crawler/multitask/concurrent/model"
)

//const(
	// ageRe  = `<div class="m-btn purple" [^>]*>([\d]+)岁</div>`
	// marriageRe = `<div class="m-btn purple" [^>]*>([^<]+)</div>`
//)

// 预先编译
// <div class="m-btn purple" data-v-bff6f798="">25岁</div>
var ageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)岁</div>`)
// <div class="m-btn purple" data-v-bff6f798="">162cm</div>
var heightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)cm</div>`)
// <div class="m-btn purple" data-v-bff6f798="">53kg</div>
var weightRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([\d]+)kg</div>`)

// <div class="m-btn purple" data-v-bff6f798="">未婚</div>
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
// <div class="m-btn purple" data-v-bff6f798="">月收入:3-5千</div>
var incomeRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>月收入:([^<]+)</div>`)
// <div class="m-btn purple" data-v-bff6f798="">会计</div>
var occupationRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
// <div class="m-btn purple" data-v-bff6f798="">大学本科</div>
var educationRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
// <div class="m-btn purple" data-v-bff6f798="">天秤座(09.23-10.22)</div>
var xingzuoRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>([^<]+)</div>`)
// <div class="m-btn pink" data-v-bff6f798="">籍贯:湖北武汉</div>
var hukouRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>籍贯:([^<]+)</div>`)
// <div class="m-btn pink" data-v-bff6f798="">租房</div>
var houseRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>([^<]+)</div>`)
// <div class="m-btn pink" data-v-bff6f798="">未买车</div>
var carRe = regexp.MustCompile(`<div class="m-btn pink" [^>]*>([^<]+)</div>`)

// 用户名 改为由城市解析器ParseCity传入
// <h1 class="nickName" data-v-5b109fc3="">等风也等你</h1>
// var nameRe = regexp.MustCompile(`<h1 class="nickName" [^>]*>([^<]+)</h1>`)

func ParseProfile(contents []byte,name string) engine.ParseResult{
	profile := model.Profile{}
	profile.Name = name

	age,err := strconv.Atoi(extractString(contents,ageRe))
	if err == nil{
		profile.Age = age
	}

	height,err := strconv.Atoi(extractString(contents,heightRe))
	if err == nil{
		profile.Height = height
	}

	weight,err := strconv.Atoi(extractString(contents,weightRe))
	if err == nil{
		profile.Weight = weight
	}

	profile.Marriage = extractString(contents,marriageRe)
	profile.Income = extractString(contents,incomeRe)
	profile.Hukou = extractString(contents,hukouRe)


	//profile.Xingzuo = extractString(contents,xingzuoRe)
	//profile.Education = extractString(contents,educationRe)
	//profile.Occupation = extractString(contents,occupationRe)
	//profile.House = extractString(contents,houseRe)
	//profile.Car = extractString(contents,carRe)

	result := engine.ParseResult{
		Items:[]interface{}{profile},
	}
	return result
}

func extractString(contents []byte,re *regexp.Regexp)string{
	match := re.FindSubmatch(contents)
	if len(match) >= 2{
		// match[0] 匹配到的字符串本身
		// match[1] 需要()语法提取的第1个值
		return string(match[1])
	}else{
		return ""
	}
}