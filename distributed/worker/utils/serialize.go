package utils

import (
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"github.com/georgefzc/crawler/config"
)

//engine.Parser can`t be transform at JsonRPC,so need Serialize and Deserialize result.
type Request struct {
	Url        string
	ParserName string
}
type Item struct {
	Url     string
	Id      string
	Payload interface{}
}
type Result struct {
	Requests []Request
	Items    []parser.Item
}

func SerializeRequest(request parser.Request) Request {
	var name string
	switch request.Parser.(type) {
	case *parser.CityList:
		name = config.ParseCityListName
	case *parser.City:
		name = config.ParseCityName
	case *parser.Person:
		name = config.ParsePersonName
	}
	return Request{
		Url:        request.Url,
		ParserName: name,
	}
}
func DeserializeRequest(sReq Request) parser.Request {
	request := parser.Request{
		Url: sReq.Url,
	}
	switch sReq.ParserName {
	case config.ParseCityListName:
		request.Parser = &parser.CityList{}
	case config.ParseCityName:
		request.Parser = &parser.City{}
	case config.ParsePersonName:
		request.Parser = &parser.Person{}
	}

	return request
}
func SerializeResult(result parser.Result) Result {
	sRes := Result{
		Items: result.Items,
	}
	for _, r := range result.Requests {
		sRes.Requests = append(sRes.Requests, SerializeRequest(r))
	}

	return sRes
}
func DeserializeResult(sRes Result) parser.Result {
	result := parser.Result{
		Items: sRes.Items,
	}
	for _, r := range sRes.Requests {
		result.Requests = append(result.Requests, DeserializeRequest(r))
	}

	return result
}
