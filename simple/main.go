package main

import (
	"github.com/georgefzc/crawler/simple/engine"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
)

const cityListURL = "http://www.zhenai.com/zhenghun"

//Start Engine.
func main() {

	//e := engine.SerialEngine{}
	//e.Run(parser.Request{
	//	Url:    cityListURL,
	//	Parser: &parser.CityList{},
	//})

	e := engine.ConcurrentEngine{}
	e.Run(parser.Request{
		Url:    cityListURL,
		Parser: &parser.CityList{},
	})

}
