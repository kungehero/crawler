package main

import (
	"crawler-go/engine"
	"crawler-go/friendswebsite/parser"
)

func main() {
	e := engine.ConcurrentRead{}
	e.Run(engine.Request{Url: `http://www.zhenai.com/jiaoyou`, ParserFunc: parser.ParserCityUrl})

	//e.Run(engine.Request{Url: `http://www.zhenai.com/jiaoyou`, ParserFunc: parser.ParserProfileMessage})
}
