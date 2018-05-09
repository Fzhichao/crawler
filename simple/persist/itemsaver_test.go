package persist

import (
	"testing"
	"github.com/georgefzc/crawler/simple/zhenai/parser"
	"gopkg.in/olivere/elastic.v5"
	"github.com/georgefzc/crawler/config"
)

func TestSaver(t *testing.T) {
	expected := parser.Item{
		Url: "http://album.zhenai.com/u/108906739",
		Id:  "108906739",
		Payload: parser.Person{
			Age:    "34岁",
			Height: "162CM",
			Weight: "57KG",
			Salary: "3001-5000元",
			Gender: "女",
			Name:   "安静的雪",
			House:  "已购房",
			Car:    "未购车",
		},
	}
	client, err := elastic.NewClient(elastic.SetURL(config.ElasticWinUrl), elastic.SetSniff(false))
	if err != nil {
		t.Errorf("elastic start err: %v", err)
	}
	err = Save(client, "data_test", "zhenai", expected)
	if err != nil {
		t.Errorf("ItemSaver failed: %v", err)
	}
	//The test is not completely...
}
