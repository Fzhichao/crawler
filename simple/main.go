package main

import (
	"github.com/georgefzc/crawler/simple/engine"
	"github.com/georgefzc/crawler/simple/persist"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"github.com/georgefzc/crawler/config"
	"log"
)

//Start Engine.
func main() {

	itemChan, err := persist.ItemSaver("data", "zhenai")
	if err != nil {
		log.Printf("ItemSaver start error: %v", err)
		return
	}
	e := engine.ConcurrentEngine{
		WorkerNum: config.WorkerNumber,
		ItemChan:  itemChan,
	}
	e.Run(parser.Request{
		Url:    config.SeedCityListURL,
		Parser: &parser.CityList{},
	})

}
