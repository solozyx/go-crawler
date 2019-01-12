package main

import (
	"regexp"
	"fmt"
)

// 正则表达式 更加文本规则匹配字符串
const(
	text = `my email is zhaoyingxin@aliyun.com@baidu.com
			your email is abc@def.me
			email is xxx@yyy.org 官方
			email is xxx@yyy.com.cn 官方中国站 
			`
)

func main() {
	// 正则表达式匹配器
	// 语法错误返回err
	// re,err := regexp.Compile("zhaoyingxin@aliyun.com")

	// 参数必须符合正则表达式语法
	// emailReg := `[a-zA-Z0-9]+@[a-zA-Z0-9]+\.[a-zA-Z0-9.]+`
	emailReg := `([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`
	re := regexp.MustCompile(emailReg)
	// match := re.FindString(text)
	// match := re.FindAllString(text,-1)
	match := re.FindAllStringSubmatch(text,-1)
	for _,m := range match{
		fmt.Println(m)
	}
}