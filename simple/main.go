package main

import (
	"github.com/georgefzc/crawler/simple/engine"
	"github.com/georgefzc/crawler/simple/zhenai/parse"
)

const cityListURL = "http://www.zhenai.com/zhenghun"

//This is a simple single thread main
func main() {
	e := engine.SerialEngine{}
	e.Run(engine.Request{
		Url:    cityListURL,
		Parser: &parse.CityList{},
	})
}
