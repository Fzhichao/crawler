package engine

type Request struct {
	Url    string
	Parser Parser
}

type Item struct {
	Payload interface{}
}

type ParseResult struct {
	Requests []Request
	Items    []Item
}

type Parser interface {
	Parse(contents []byte) *ParseResult
}
