package parser

import (
	"regexp"
)

//It is for precompile
var (
	nameRe   = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1> `)
	ageRe    = regexp.MustCompile(`<td><span class="label">年龄：</span>([^<]+)</td>`)
	salaryRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	heightRe = regexp.MustCompile(` <td><span class="label">身高：</span><span field="">([^<]+)</span></td>`)
	weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
	houseRe  = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe    = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
)
var guessRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)

//Person implements the Parse interface
type Person struct {
	Name   string
	Age    string
	Salary string
	Gender string
	Height string
	Weight string
	House  string
	Car    string
}

//Parse CityList html contents
func (p *Person) Parse(contents []byte) *Result {
	person := Person{
		Name:   extractString(contents, nameRe),
		Age:    extractString(contents, ageRe),
		Salary: extractString(contents, salaryRe),
		Gender: extractString(contents, genderRe),
		Height: extractString(contents, heightRe),
		Weight: extractString(contents, weightRe),
		House:  extractString(contents, houseRe),
		Car:    extractString(contents, carRe),
	}

	res := Result{}
	guessMatches := guessRe.FindAllSubmatch(contents, -1)
	for _, m := range guessMatches {
		res.Requests = append(res.Requests, Request{
			Url:    string(m[1]),
			Parser: &Person{},
		})
	}
	res.Items = append(res.Items, Item{
		Payload: person,
	})

	return &res
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) > 1 {
		return string(match[1])
	} else {
		return ""
	}
}
