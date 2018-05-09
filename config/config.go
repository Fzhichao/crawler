package config

const (
	ItemSaverRpcHost = ":1234"                      //	ItemSaver Server IP:Port
	ElasticWinUrl    = "http://192.168.99.100:9200" //	ElasticSearch IP:Port
	ElasticIndex     = "data"
	ElasticType      = "zhenai"
)

const (
	SeedCityListURL = "http://www.zhenai.com/zhenghun"
	//Number of worker goroutine
	WorkerNumber    = 100
)
