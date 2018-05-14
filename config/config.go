package config

const (
	ElasticWinUrl = "http://192.168.99.100:9200" //	ElasticSearch server default IP:Port
	ElasticIndex  = "data"
	ElasticType   = "zhenai"
)

const (
	ItemSaverServiceRpc = "ItemSaverService.Save" //Service.method
	CrawlerServiceRpc   = "CrawlerService.Crawl"  //Service.method
)

const (
	SeedCityListURL   = "http://www.zhenai.com/zhenghun"
	WorkerNumber      = 100 //Number of worker goroutine
	ParseCityListName = "cityList"
	ParseCityName     = "city"
	ParsePersonName   = "person"
)
