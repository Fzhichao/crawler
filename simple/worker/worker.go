package worker

import (
	"github.com/georgefzc/crawler/simple/fetcher"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"github.com/georgefzc/crawler/simple/scheduler"
)

type WorkProcess func(r parser.Request) (*parser.Result, error)

//Worker maybe type a struct,at there through function.
//Golang is pretty convenient
func Run(in chan parser.Request, out chan<- *parser.Result, s *scheduler.Scheduler, work WorkProcess) {
	go func() {
		for {
			s.SubmitWorker(in)
			request := <-in
			res, err := work(request)
			if err != nil {
				//log.Printf("WorkerErr: Fetching Url %s: %v", request.Url, err)
				continue
			}
			out <- res
		}
	}()
}

func Work(r parser.Request) (*parser.Result, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		return nil, err
	}

	return r.Parser.Parse(body, r.Url)
}

func NewWorker() chan parser.Request {
	return make(chan parser.Request)
}
