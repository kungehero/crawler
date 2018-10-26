package engine

import (
	"crawler-go/fetcher"
	"fmt"
	"log"
)

type ConcurrentRead struct {
}

func (Con ConcurrentRead) Run(seeds ...Request) { //engine.Request{Url: `http://www.zhenai.com/jiaoyou`, ParserFunc: parser.ParserProfileMessage}
	var requests []Request
	var items []interface{}
	for _, s := range seeds {
		requests = append(requests, s)
	}
	if len(requests) > 0 {
		re := requests[0]
		requests = requests[1:]
		body, err := fetcher.FetchUrlData(re.Url)
		if err != nil {
			log.Printf("Fetch Url is %s", re.Url)
		}
		parserresult := re.ParserFunc(body)
		requests = append(requests, parserresult.Requests...) //拼接切片

		for _, mm := range requests {
			body, _ := fetcher.FetchUrlData(mm.Url)
			requests = append(requests, mm.ParserFunc(body).Requests...) //拼接切片

		}

		for _, s := range requests {
			body, _ := fetcher.FetchUrlData(s.Url)
			items = append(items, s.ParserFunc(body).Items...) //拼接切片

		}

		for _, item := range items {
			fmt.Printf("%s\n", item)
		}
	}
}
