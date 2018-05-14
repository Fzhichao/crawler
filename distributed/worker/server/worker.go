package main

import (
	"github.com/georgefzc/crawler/distributed/worker"
	"github.com/georgefzc/crawler/distributed/rpccommon"
	"log"
	"flag"
	"fmt"
)

var port = flag.Int("port", 0, "Worker server port to listen on")

//Start a server listen worker client send request to fetch and parse.
func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("Must specify a port!")
		return
	}

	log.Fatal(rpccommon.ServerRPC(fmt.Sprintf(":%d", *port), &worker.CrawlerService{}))
}
