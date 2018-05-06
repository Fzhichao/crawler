package engine

import (
	"github.com/georgefzc/crawler/simple/fetcher"
	"log"
	"errors"
)

func Worker(r Request) (*ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("WorkFetcher: error fetching url %s: %v", r.Url, err)
		return &ParseResult{}, err
	}
	if body == nil {
		return &ParseResult{}, errors.New("nil body return")
	}

	return r.Parser.Parse(body), nil
}
