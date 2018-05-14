package engine

import (
	"github.com/georgefzc/crawler/simple/worker"
	"github.com/georgefzc/crawler/simple/scheduler"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
)

type ConcurrentEngine struct {
	Scheduler   scheduler.Scheduler
	ItemChan    chan parser.Item
	WorkerNum   int
	WorkProcess worker.WorkProcess
}

// Engine start scheduler and worker.
// Receive parsed requests and items.
func (e *ConcurrentEngine) Run(seeds ...parser.Request) {
	e.Scheduler.Run()

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.SubmitRequest(r)
	}
	workRes := make(chan *parser.Result)
	for i := 0; i < e.WorkerNum; i++ {
		//Every worker will has a self chan.
		worker.Run(worker.NewWorker(), workRes, &e.Scheduler, e.WorkProcess)
	}

	for {
		receive := <-workRes
		for _, item := range receive.Items {
			// go func() will not block Engine.
			go func() { e.ItemChan <- item }()
		}
		for _, r := range receive.Requests {
			if isDuplicate(r.Url) {
				//log.Printf("Duplicate request: %s", r.Url)
				continue
			}
			e.Scheduler.SubmitRequest(r)
		}
	}
}

// Should use Redis or other K-V database.
var visitedUrl = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}
