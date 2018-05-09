package engine

import (
	"log"
	"github.com/georgefzc/crawler/simple/worker"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
)

type SerialEngine struct{}

//It is a engine for simple scheduling.
func (e SerialEngine) Run(seeds ...parser.Request) {
	var requestsQ []parser.Request
	for _, r := range seeds {
		requestsQ = append(requestsQ, r)
	}
	for len(requestsQ) > 0 {
		r := requestsQ[0]
		requestsQ = requestsQ[1:]

		workRes, err := worker.Work(r)
		if err != nil {
			continue
		}
		requestsQ = append(requestsQ, workRes.Requests...)

		for i, item := range workRes.Items {
			log.Printf("Got item %d: %s", i, item)
		}
	}
}
