package client

import (
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"github.com/georgefzc/crawler/distributed/rpcutils"
	"log"
)

func ItemSaver(host string) (chan parser.Item, error) {
	client, err := rpcutils.NewClient(host)
	if err != nil {
		return nil, err
	}

	out := make(chan parser.Item)
	go func() {
		count := 0
		for {
			item := <-out
			log.Printf("ItemSaver: Got item%d: %s", count, item)
			count++
			rpcRes := false
			if err := client.Call("ItemSaverService.Save", item, &rpcRes); err != nil {
				log.Printf("ItemSaverErr: save item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}
