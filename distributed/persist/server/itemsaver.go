package main

import (
	"github.com/georgefzc/crawler/distributed/persist"
	"gopkg.in/olivere/elastic.v5"
	"github.com/georgefzc/crawler/distributed/rpcutils"
	"github.com/georgefzc/crawler/config"
)

func main() {
	//TODO: Start elasticSearch
	client, err := elastic.NewClient(
		elastic.SetSniff(false),
		elastic.SetURL(config.ElasticWinUrl,
		))
	if err != nil {
		panic(err)
	}

	rpcutils.ServerRPC(config.ItemSaverRpcHost, &persist.ItemSaverService{
		Client: client,
		Index:  config.ElasticIndex,
		Typ:    config.ElasticType,
	})

}
