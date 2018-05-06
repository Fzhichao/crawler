package parse

import (
	"regexp"
	"github.com/georgefzc/crawler/simple/engine"
)

var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)

type CityList struct{}

//Parse CityList html contents
func (cl *CityList) Parse(contents []byte) *engine.ParseResult {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	res := &engine.ParseResult{}
	count := 0
	for _, m := range matches {
		count++
		if count > 2 {
			break
		}
		res.Requests = append(res.Requests, engine.Request{
			Url:    string(m[1]),
			Parser: &City{},
		})
		res.Items = append(res.Items, engine.Item{
			Payload: m[2],
		})
	}
	return res

}
