package client

import (
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"github.com/georgefzc/crawler/distributed/rpccommon"
	"log"
	"github.com/georgefzc/crawler/config"
)

func ItemSaver(host string) (chan parser.Item, error) {
	client, err := rpccommon.NewClient(host)
	if err != nil {
		return nil, err
	}
	log.Println("Connect to ItemSaver RPCServer")
	out := make(chan parser.Item)
	go func() {
		for {
			item := <-out
			rpcRes := false
			if err := client.Call(config.ItemSaverServiceRpc, item, &rpcRes); err != nil {
				log.Printf("ItemSaverErr: save item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
