package parser

import (
	"regexp"
)
//It is for precompile
var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)

//CityList implements the Parse interface
type CityList struct{}

//Parse CityList html contents
func (cl *CityList) Parse(contents []byte) *Result {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	res := &Result{}
	count := 0
	for _, m := range matches {
		count ++
		if count > 5 {
			break
		}
		res.Requests = append(res.Requests, Request{
			Url:    string(m[1]),
			Parser: &City{},
		})
		res.Items = append(res.Items, Item{
			Payload: m[2],
		})
	}
	return res

}
