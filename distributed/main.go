package main

import (
	"log"
	"github.com/georgefzc/crawler/config"
	"github.com/georgefzc/crawler/simple/engine"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	saverclient "github.com/georgefzc/crawler/distributed/persist/client"
	workerclient "github.com/georgefzc/crawler/distributed/worker/client"
	"net/rpc"
	"github.com/georgefzc/crawler/distributed/rpccommon"
	"errors"
	"flag"
	"strings"
)

var (
	itemSaverRpcHost = flag.String("ItemSaver_host", "", "ItemSaver host to connect")
	workerRpcHost    = flag.String("Worker_host", "", "Worker host to connect (comma separated)")
)

func main() {
	flag.Parse()

	itemChan, err := saverclient.ItemSaver(*itemSaverRpcHost)
	if err != nil {
		//ItemSaver must connect or crawler isn`t useful.
		panic(err)
	}

	c, err := createWorkerClientPool(strings.Split(*workerRpcHost, ","))
	if err != nil {
		//If Error,indicate no client can use.
		panic(err)
	}

	e := engine.ConcurrentEngine{
		WorkerNum:   config.WorkerNumber,
		ItemChan:    itemChan,
		WorkProcess: workerclient.Process(c),
	}
	e.Run(parser.Request{
		Url:    config.SeedCityListURL,
		Parser: &parser.CityList{},
	})
}
func createWorkerClientPool(hosts []string) (chan *rpc.Client, error) {
	var clients []*rpc.Client
	for _, h := range hosts {
		client, err := rpccommon.NewClient(h)
		if err != nil {
			log.Printf("Connect to Worker RPCServer %s error : %v", h, err)
			continue
		}
		log.Printf("Connect to Worker RPCServer %s", h)
		clients = append(clients, client)
	}
	if len(clients) == 0 {
		return nil, errors.New("no connections Worker RPCServer available")
	}
	out := make(chan *rpc.Client)
	go func() {
		for {
			for _, c := range clients {
				out <- c
			}
		}
	}()
	return out, nil
}
