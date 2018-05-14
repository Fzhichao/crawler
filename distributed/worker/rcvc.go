package worker

import (
	"github.com/georgefzc/crawler/simple/worker"
	"github.com/georgefzc/crawler/distributed/worker/utils"
	"log"
)

//JsonRPC service.method
type CrawlerService struct {
}

func (s *CrawlerService) Crawl(sReq utils.Request, sRes *utils.Result) (err error) {
	request := utils.DeserializeRequest(sReq)
	result, err := worker.Work(request)
	if err != nil {
		log.Printf("Crawler Process err : %v", err)
		return err
	}
	*sRes = utils.SerializeResult(*result)

	return nil
}
