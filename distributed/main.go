package main

import (
	"log"
	"github.com/georgefzc/crawler/config"
	"github.com/georgefzc/crawler/simple/engine"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"github.com/georgefzc/crawler/distributed/persist/client"
)

func main() {
	itemChan, err := client.ItemSaver(config.ItemSaverRpcHost)
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
