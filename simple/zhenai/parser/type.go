package parser

type Request struct {
	Url    string
	Parser Parser
}

type Item struct {
	Payload interface{}
}

type Result struct {
	Requests []Request
	Items    []Item
}

type Parser interface {
	Parse(contents []byte) *Result
}


