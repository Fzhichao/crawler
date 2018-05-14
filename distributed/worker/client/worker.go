package client

import (
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"github.com/georgefzc/crawler/config"
	"github.com/georgefzc/crawler/distributed/worker/utils"
	"github.com/georgefzc/crawler/simple/worker"
	"net/rpc"
)

//Worker client to di
func Process(clientChan chan *rpc.Client) worker.WorkProcess {

	return func(request parser.Request) (*parser.Result, error) {
		sReq := utils.SerializeRequest(request)
		var sRes utils.Result
		client := <-clientChan
		err := client.Call(config.CrawlerServiceRpc, sReq, &sRes)
		if err != nil {
			return nil, err
		}
		result := utils.DeserializeResult(sRes)

		return &result, err
	}
}
