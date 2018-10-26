package parser

import (
	"crawler-go/engine"
	"regexp"
)

func ParserCity(urls []byte) engine.ParserResult {
	re := regexp.MustCompile(`<a class="name" href="(http://album.zhenai.com/u/[0-9]+)"[^>]+>([^<]+)</a>`)

	addr := re.FindAllSubmatch(urls, -1)
	result := engine.ParserResult{}

	for _, m := range addr {
		result.Requests = append(result.Requests, engine.Request{Url: string(m[1]), ParserFunc: ParserProfileMessage})
		result.Items = append(result.Items, m[2])
	}
	return result
}
