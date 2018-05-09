package parser

import (
	"regexp"
)

//It is for precompile
var (
	idRe     = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)
	nameRe   = regexp.MustCompile(`<h1 class="ceiling-name ib fl fs24 lh32 blue">([^<]+)</h1> `)
	ageRe    = regexp.MustCompile(`<td><span class="label">年龄：</span>([^<]+)</td>`)
	salaryRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
	genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
	heightRe = regexp.MustCompile(` <td><span class="label">身高：</span><span field="">([^<]+)</span></td>`)
	weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([^<]+)</span></td>`)
	houseRe  = regexp.MustCompile(`<td><span class="label">住房条件：</span><span field="">([^<]+)</span></td>`)
	carRe    = regexp.MustCompile(`<td><span class="label">是否购车：</span><span field="">([^<]+)</span></td>`)
)
var guessPersonRe = regexp.MustCompile(`<a class="exp-user-name"[^>]*href="(http://album.zhenai.com/u/[\d]+)">([^<]+)</a>`)

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
func (p *Person) Parse(contents []byte, url string) (*Result, error) {
	person := extractPerson(contents)
	res := Result{}
	matches := guessPersonRe.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		res.Requests = append(res.Requests, Request{
			Url:    string(m[1]),
			Parser: &Person{},
		})
	}
	res.Items = append(res.Items, Item{
		Url:     url,
		Id:      extractString([]byte(url), idRe),
		Payload: person,
	})

	return &res, nil
}
func extractPerson(contents []byte) Person {
	p := Person{}
	if name := nameRe.FindSubmatch(contents); len(name) > 1 {
		p.Name = string(name[1])
	}
	if age := ageRe.FindSubmatch(contents); len(age) > 1 {
		p.Age = string(age[1])
	}
	if salary := salaryRe.FindSubmatch(contents); len(salary) > 1 {
		p.Salary = string(salary[1])
	}
	if gender := genderRe.FindSubmatch(contents); len(gender) > 1 {
		p.Gender = string(gender[1])
	}
	if height := heightRe.FindSubmatch(contents); len(height) > 1 {
		p.Height = string(height[1])
	}
	if weight := weightRe.FindSubmatch(contents); len(weight) > 1 {
		p.Weight = string(weight[1])
	}
	if house := houseRe.FindSubmatch(contents); len(house) > 1 {
		p.House = string(house[1])
	}
	if car := carRe.FindSubmatch(contents); len(car) > 1 {
		p.Car = string(car[1])
	}
	return p
}
func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) > 1 {
		return string(match[1])
	} else {
		return ""
	}
}
