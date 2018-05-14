package main

import (
	"log"
	"github.com/georgefzc/crawler/simple/engine"
	"github.com/georgefzc/crawler/simple/persist"
	"github.com/georgefzc/crawler/config"
	"github.com/georgefzc/crawler/simple/worker"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
)

//Start Engine.
func main() {

	itemChan, err := persist.ItemSaver(config.ElasticIndex, config.ElasticType)
	if err != nil {
		log.Printf("ItemSaver start error: %v", err)
		return
	}
	e := engine.ConcurrentEngine{
		WorkerNum:   config.WorkerNumber,
		ItemChan:    itemChan,
		WorkProcess: worker.Work,
	}
	e.Run(parser.Request{
		Url:    config.SeedCityListURL,
		Parser: &parser.CityList{},
	})

}
