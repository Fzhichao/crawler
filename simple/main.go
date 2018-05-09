package main

import (
	"github.com/georgefzc/crawler/simple/engine"
	"github.com/georgefzc/crawler/simple/persist"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"log"
)

const cityListURL = "http://www.zhenai.com/zhenghun"
const workerNumber = 100

//Start Engine.
func main() {

	itemChan, err := persist.ItemSaver("data", "zhenai")
	if err != nil {
		log.Printf("ItemSaver start error: %v", err)
		return
	}
	e := engine.ConcurrentEngine{
		WorkerNum: workerNumber,
		ItemChan:  itemChan,
	}
	e.Run(parser.Request{
		Url:    cityListURL,
		Parser: &parser.CityList{},
	})

}
