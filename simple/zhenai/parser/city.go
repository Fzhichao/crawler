package parser

import (
	"regexp"
)

//It is for precompile
var (
	personRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
)

//City implements the Parse interface
type City struct{}

//Parse City html contents
func (c *City) Parse(contents []byte) *Result {
	matches := personRe.FindAllSubmatch(contents, -1)

	res := &Result{}
	for _, m := range matches {
		res.Requests = append(res.Requests, Request{
			Url:    string(m[1]),
			Parser: &Person{},
		})
		res.Items = append(res.Items, Item{
			Payload: m[2],
		})
	}
	return res
}
