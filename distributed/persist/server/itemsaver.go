package main

import (
	"github.com/georgefzc/crawler/distributed/persist"
	"gopkg.in/olivere/elastic.v5"
	"github.com/georgefzc/crawler/distributed/rpccommon"
	"github.com/georgefzc/crawler/config"
	"log"
	"flag"
	"fmt"
)

var port = flag.Int("port", 0, "ItemSaver server port to listen")

//Start elasticSearch client to connect.
//Start server to listen ItemSaver client send person data to save.
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port")
		return
	}

	//TODO: Start elasticSearch
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.ElasticWinUrl,
		))
	if err != nil {
		//elasticSearch must connect or crawler isn`t useful.
		panic(err)
	}
	log.Println("ItemSaver connect to elasticSearch Server")

	log.Fatal(rpccommon.ServerRPC(fmt.Sprintf(":%d", *port), &persist.ItemSaverService{
		Client: client,
		Index:  config.ElasticIndex,
		Typ:    config.ElasticType,
	}))

}
