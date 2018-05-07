package fetcher

import (
	"bufio"
	"log"
	"time"
	"net/http"
	"io/ioutil"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

var rateLimiter = time.Tick(time.Second / 20)

// Get URL and return UTF8 contents
func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("%s Fetch StatusCode is %s, not ok\n", url, resp.Status)
		return nil, err
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	newReader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(newReader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Println("Error: determineEncoding Peek err ,return default UTF8")
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
