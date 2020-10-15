package spider_api

import (
	"context"
	"encoding/json"
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"net/http"

	"perch/web/metric"
	"perch/web/model"
)

type pageInfo struct {
	StatusCode int
	Links      map[string]int
}

func CreateCollySpider(writer http.ResponseWriter, request *http.Request) {
	metric.ProcessMetricFunc(writer, request, nil, func(ctx context.Context, bean interface{}, respone *model.ResultReponse) error {

		URL := request.URL.Query().Get("url")
		if URL == "" {
			log.Println("missing URL argument")

		}
		log.Println("visiting", URL)

		c := colly.NewCollector()

		p := &pageInfo{Links: make(map[string]int)}

		// count links
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Request.AbsoluteURL(e.Attr("href"))
			if link != "" {
				p.Links[link]++
			}
		})

		// extract status code
		c.OnResponse(func(r *colly.Response) {
			log.Println("response received", r.StatusCode)
			p.StatusCode = r.StatusCode
		})
		c.OnError(func(r *colly.Response, err error) {
			log.Println("error:", r.StatusCode, err)
			p.StatusCode = r.StatusCode
		})

		c.Visit(URL)

		// dump results
		b, err := json.Marshal(p)
		if err != nil {
			log.Println("failed to serialize response:", err)

		}

		respone.Spec = b
		respone.Kind = "blockchain"
		respone.Code = http.StatusOK

		return nil
	})

}
