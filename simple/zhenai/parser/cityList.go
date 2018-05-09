package parser

import (
	"regexp"
)

//It is for precompile
var cityListRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`)

//CityList implements the Parse interface
type CityList struct{}

//Parse CityList html contents
func (cl *CityList) Parse(contents []byte, _ string) (*Result, error) {
	matches := cityListRe.FindAllSubmatch(contents, -1)
	res := &Result{}
	limit := 0	//just temp
	for _, m := range matches {
		if limit++; limit > 10 {
			break
		}
		res.Requests = append(res.Requests, Request{
			Url:    string(m[1]),
			Parser: &City{},
		})
		//I don`t need city item ,if needed item can be seeded
		//log.Printf("%s",m[2])
		//res.Items = append(res.Items, Item{
		//	Payload: m[2],
		//})
	}

	return res, nil
}
