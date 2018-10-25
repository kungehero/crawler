package fetcher

import (
	"io/ioutil"
	"net/http"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
)

func FetchUrlData(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		panic("网络连接失败！")
	}
	defer resp.Body.Close()

	utf8 := transform.NewReader(resp.Body, simplifiedchinese.GBK.NewDecoder())
	all, _ := ioutil.ReadAll(utf8)
	return all
}
