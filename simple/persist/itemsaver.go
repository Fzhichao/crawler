package persist

import (
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"gopkg.in/olivere/elastic.v5"
	"context"
	"log"
	"errors"
)

func ItemSaver(index string, typ string) (chan parser.Item, error) {
	client, err := elastic.NewClient(elastic.SetURL(userUrl), elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}
	out := make(chan parser.Item)
	go func() {
		count := 0
		for {
			item := <-out
			count++
			log.Printf("ItemSaver: Got item%d: %s", count, item)
			if err := save(client, index, typ, item); err != nil {
				log.Printf("ItemSaverErr: save item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

const userUrl = "http://192.168.99.100:9200"

func save(client *elastic.Client, index string, typ string, item parser.Item) error {
	if index == "" || typ == "" {
		return errors.New("must appoint index and typ")
	}
	indexService := client.Index().Index(index).Type(typ).BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err := indexService.Do(context.Background())

	return err

}
