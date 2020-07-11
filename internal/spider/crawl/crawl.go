package crawl

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type SpiderCrawer interface {
	CrawlerFetch(url string) error
}

func CrawlerFetch(url string, method string, depth int, requestData io.Reader, ctx *context.Context, requestHeader http.Header, checkRevisited bool) error {
	var err error
	httpClient := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return err
	}
	req.Header.Set("User-Agent", "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1)")
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Printf("request url %s,response is %s\n", url, string(body))

	return nil
}
