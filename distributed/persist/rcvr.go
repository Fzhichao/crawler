package persist

import (
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"github.com/georgefzc/crawler/simple/persist"
	"gopkg.in/olivere/elastic.v5"
	"log"
)

type ItemSaverService struct {
	Client     *elastic.Client
	Index, Typ string
}

func (s *ItemSaverService) Save(item parser.Item, result *bool) error {
	err := persist.Save(s.Client, s.Index, s.Typ, item)
	log.Printf("Save Item %v", item)
	if err == nil {
		*result = true
	} else {
		*result = false
	}
	return err
}
