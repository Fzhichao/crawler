package parser

import (
	"testing"
	"io/ioutil"
)

func TestCityList_Parse(t *testing.T) {
	const resultSize = 470
	expected := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}

	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}
	c := CityList{}
	res, _ := c.Parse(contents, "")

	if len(res.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(res.Requests))
	}
	for i, url := range expected {
		if res.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, res.Requests[i].Url)
		}
	}

}
func TestPerson_Parse(t *testing.T) {
	expected := Item{
		Url: "http://album.zhenai.com/u/108906739",
		Id:  "108906739",
		Payload: Person{
			Name:   "安静的雪",
			Age:    "34岁",
			Salary: "3001-5000元",
			Gender: "女",
			Height: "162CM",
			Weight: "57KG",
			House:  "已购房",
			Car:    "未购车",
		},
	}
	contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}
	p := Person{}
	res, _ := p.Parse(contents, "http://album.zhenai.com/u/108906739")
	actual := res.Items[0]

	if actual != expected {
		t.Errorf("expected %v; but was %v", expected, actual)
	}

}
