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
	metric.ProcessMetricFunc(writer, request, nil)

}
