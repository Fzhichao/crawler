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
	"fmt"
)

//Server Anti-Spider may limit IP.
var rateLimiter = time.Tick(time.Second / 20)

// Get URL and return UTF8 contents
func Fetch(url string) ([]byte, error) {
	<-rateLimiter
	log.Printf("Fetching Url %s", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Fetch StatusCode %s\n", resp.Status)
	}
	bodyReader := bufio.NewReader(resp.Body)
	e := determineEncoding(bodyReader)
	newReader := transform.NewReader(resp.Body, e.NewDecoder())
	return ioutil.ReadAll(newReader)
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	bytes, err := r.Peek(1024)
	if err != nil {
		log.Println("DetermineEncoding Peek err ,return default UTF8")
		return unicode.UTF8
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")

	return e
}
