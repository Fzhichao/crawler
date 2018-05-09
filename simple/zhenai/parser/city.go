package parser

import (
	"regexp"
)

//It is for precompile
var (
	personRe   = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)
	moreCityRe = regexp.MustCompile(`(http://www.zhenai.com/zhenghun/[^"]+)`)
)

//City implements the Parse interface
type City struct{}

//Parse City html contents
func (c *City) Parse(contents []byte, _ string) (*Result, error) {
	matches := personRe.FindAllSubmatch(contents, -1)
	res := &Result{}
	for _, m := range matches {
		res.Requests = append(res.Requests, Request{
			Url:    string(m[1]),
			Parser: &Person{},
		})
		//I don`t need person name item ,if needed item can be seeded
		//log.Printf("%s",m[2])
		//res.Items = append(res.Items, Item{
		//	Payload: m[2],
		//})
	}

	matches = moreCityRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		res.Requests = append(res.Requests, Request{
			Url:    string(m[1]),
			Parser: &City{},
		})
	}
	return res, nil
}
