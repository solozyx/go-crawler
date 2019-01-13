package fetcher

import (
	"net/http"
	"fmt"
	"bufio"
	"io/ioutil"
	"time"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
)

// TODO 待爬取目标网站如果爬取网络流量正常稳定可以适当减少等待时间

// 10毫秒执行一次请求
var rateLimiter = time.Tick(10 * time.Millisecond)

// fetch到的网页数据 该url不能获取数据则err
func Fetch(url string) ([]byte, error) {
	<- rateLimiter

	// resp,err := http.Get(url)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Set("User-Agent",
		`Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36`)
	var httpClient = http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,fmt.Errorf("wrong status code = %d",resp.StatusCode)
	}

	// e := determineEncoding(resp.Body)
	// utf8Reader := transform.NewReader(resp.Body,e.NewDecoder())

	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	utf8Reader := transform.NewReader(bodyReader,e.NewDecoder())
	return ioutil.ReadAll(utf8Reader)
}

// func determineEncoding(r io.Reader) encoding.Encoding{
func determineEncoding(r *bufio.Reader) encoding.Encoding{
	// bytes,err := bufio.NewReader(r).Peek(1024)
	bytes,err := r.Peek(1024)
	if err != nil{
		// Peek失败 不代表该网页文本不可读 返回默认编码
		fmt.Printf("Fetcher err = %v",err)
		return unicode.UTF8
	}
	e,_,_ := charset.DetermineEncoding(bytes,"")

	return e
}