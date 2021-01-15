/**
API
假数据生成  https://godoc.org/github.com/icrowley/fake
ref
*/
package data_api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/icrowley/fake"
	"perch/web/metric"
	"perch/web/model"
)

func GetFakeUsersHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func GetFakeEmailsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func GetFakeCrediCardHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func GetFakeIPHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func GetFakeLocHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}
func GetFakeTimesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}
