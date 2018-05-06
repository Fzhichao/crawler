package parse

import (
	"github.com/georgefzc/crawler/simple/engine"
	"regexp"
)

var (
	personRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
)

type City struct{}

//Parse City html contents
func (c *City) Parse(contents []byte) *engine.ParseResult {
	matches := personRe.FindAllSubmatch(contents, -1)

	res := &engine.ParseResult{}
	for _, m := range matches {
		res.Requests = append(res.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: &Person{},
		})
		res.Items = append(res.Items, engine.Item{
			Payload: m[2],
		})
	}
	return res
}
