package worker

import (
	"fmt"
	"log"
	"github.com/georgefzc/crawler/simple/fetcher"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"github.com/georgefzc/crawler/simple/scheduler"
)
//Worker maybe type a struct,at there through function.
//Golang is pretty convenient
func Run(in chan parser.Request, out chan<- *parser.Result, s scheduler.Scheduler) {
	go func() {
		for {
			s.SubmitWorker(in)
			request := <-in
			res, err := Work(request)
			if err != nil {
				continue
			}
			out <- res
		}
	}()
}

func Work(r parser.Request) (*parser.Result, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("WorkFetcher: error fetching url %s: %v", r.Url, err)
		return &parser.Result{}, err
	}
	if body == nil {
		return &parser.Result{}, fmt.Errorf("nil body return")
	}

	return r.Parser.Parse(body), nil
}

func NewWorker() chan parser.Request {
	return make(chan parser.Request)
}
