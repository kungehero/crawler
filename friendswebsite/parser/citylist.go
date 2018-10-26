package parser

import (
	"crawler-go/engine"
	"regexp"
)

func ParserCityUrl(urls []byte) engine.ParserResult {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/jiaoyou/[0-9a-z]+)"[^>]*>(阿[^<]+)</a>`)
	addr := re.FindAllSubmatch(urls, -1)
	result := engine.ParserResult{}

	for _, m := range addr {
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: ParserCity})
		result.Items = append(result.Items, m[2])
	}
	return result
}
