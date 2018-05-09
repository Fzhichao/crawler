package scheduler

import (
	"github.com/georgefzc/crawler/simple/zhenai/parser"
)

type worker chan parser.Request
type Scheduler struct {
	RequestChan chan parser.Request
	// We maintain a channel of worker
	// When has free worker,submit it to scheduler`s queue
	// request will be send to the worker that chan parser.Request
	WorkerChan chan worker
}

//Scheduler maintain requests and workers queue.
//Receive requests from Engine or free workers from worker, Respective queued.
//Both there are,send request to worker.
func (s *Scheduler) Run() {
	s.RequestChan = make(chan parser.Request)
	s.WorkerChan = make(chan worker)
	go func() {
		var requestsQ []parser.Request
		var workersQ []worker
		for {
			var activeRequest parser.Request
			var activeWorker worker
			if len(requestsQ) > 0 && len(workersQ) > 0 {
				activeRequest = requestsQ[0]
				activeWorker = workersQ[0]
			}
			select {
			case r := <-s.RequestChan:
				requestsQ = append(requestsQ, r)
			case w := <-s.WorkerChan:
				workersQ = append(workersQ, w)
			case activeWorker <- activeRequest:
				requestsQ = requestsQ[1:]
				workersQ = workersQ[1:]
			}
		}
	}()
}
func (s *Scheduler) SubmitWorker(w worker) {
	s.WorkerChan <- w
}
func (s *Scheduler) SubmitRequest(r parser.Request) {
	s.RequestChan <- r
}
