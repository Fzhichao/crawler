package engine

import "log"

type SerialEngine struct{}

//It is a engine for simple scheduling
func (e SerialEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		workRes, err := Worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, workRes.Requests...)

		for i, item := range workRes.Items {
			log.Printf("Got item %d: %s", i, item)
		}
	}
}
