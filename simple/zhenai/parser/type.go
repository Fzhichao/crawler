package parser

type Request struct {
	Url    string
	Parser Parser
}

type Item struct {
	Url     string
	Id      string
	Payload interface{}
}

type Result struct {
	Requests []Request
	Items    []Item
}

type Parser interface {
	Parse([]byte, string) (*Result, error)
}
