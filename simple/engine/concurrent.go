package engine

import (
	"log"
	"github.com/georgefzc/crawler/simple/worker"
	"github.com/georgefzc/crawler/simple/scheduler"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
)

const workerNumber = 100

type ConcurrentEngine struct {
	Scheduler scheduler.Scheduler
	workerNum int
}

// Engine start scheduler and worker.
// Receive parsed requests and items.
func (e *ConcurrentEngine) Run(seeds ...parser.Request) {
	e.workerNum = workerNumber
	e.Scheduler.Run()

	for _, r := range seeds {
		e.Scheduler.SubmitRequest(r)
	}
	workRes := make(chan *parser.Result)
	for i := 0; i < e.workerNum; i++ {
		//Every worker will has a self chan.
		worker.Run(worker.NewWorker(), workRes, e.Scheduler)
	}

	for {
		receive := <-workRes
		for _, item := range receive.Items {
			log.Printf("Got Item: %s\n", item)
		}
		for _, r := range receive.Requests {
			e.Scheduler.SubmitRequest(r)
		}
	}
}
