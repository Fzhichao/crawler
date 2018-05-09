package persist

import (
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"gopkg.in/olivere/elastic.v5"
	"github.com/georgefzc/crawler/config"
	"context"
	"log"
	"errors"
)

func ItemSaver(index, typ string) (chan parser.Item, error) {
	//TODO: Start elasticSearch
	client, err := elastic.NewClient(elastic.SetURL(config.ElasticWinUrl), elastic.SetSniff(false))
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
			if err := Save(client, index, typ, item); err != nil {
				log.Printf("ItemSaverErr: save item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func Save(client *elastic.Client, index, typ string, item parser.Item) error {
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
