package main

import(
	"net/http"
	"io"
	"io/ioutil"
	"fmt"
	"bufio"
	"regexp"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	// "golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/net/html/charset"
)

func main() {
	// 获取珍爱网城市列表 
	resp,err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	
	// 判断响应头
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error : status code = ",resp.StatusCode)
		return 
	}

	// TODO NOTICE <meta charset="utf-8"> 珍爱网现在编码 
	// 
	// 网页编码是gbk Go程序编码 utf-8 
	// 把1个Reader 变形为 另1个Reader 
	// 网页gbk -> Go utf-8
	// utf8Reader := transform.NewReader(resp.Body,
	// simplifiedchinese.GBK.NewDecoder())
	// resp.Body gbk Reader -> utf8Reader
	// all,err := ioutil.ReadAll(utf8Reader)

	// all,err := ioutil.ReadAll(resp.Body)
	// 
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())
	all,err := ioutil.ReadAll(utf8Reader)
	if err != nil{
		panic(err)
	}
	printCityList(all)
}

func determineEncoding(r io.Reader) encoding.Encoding{
	// 1024byte 
	// 从 r读取1024byte 后 就无法再次读取到 这 1024byte 了 
	// 用bufio装一下
	bytes,err := bufio.NewReader(r).Peek(1024)
	if err != nil{
		panic(err)
	}
	// contentType 传空 
	e,_,_ := charset.DetermineEncoding(bytes,"")
	return e 
}

// <a href="http://www.zhenai.com/zhenghun/aba" data-v-473e2ba0="">阿坝</a>
func printCityList(contents []byte){
	// regExp := `<a href="http://www.zhenai.com/zhenghun/[0-9a-z]+"[^>]*>[^<]+</a>`
	regExp := `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
	re := regexp.MustCompile(regExp)
	// matches := re.FindAll(contents,-1)
	matches := re.FindAllSubmatch(contents,-1)
	for _,m := range matches{
		// for _,subM := range m {
		//	fmt.Printf("%s  ",subM)
		// }
		// fmt.Println()
		fmt.Printf("City : %s , URL : %s \n",m[2],m[1])
	}
	fmt.Printf("Matches found citys count = %d \n",len(matches))
}