/**
文本资源 API
*/
package resources_api

import (
	"context"
	_ "fmt"
	"github.com/gorilla/mux"
	"net/http"
	database "perch/database/mysql"
	"perch/web/metric"
	"perch/web/model"
	resource "perch/web/model/resources"
)

func GetResourcesBlogsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func SpecGetResourcesBlogsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func CreateResourcesBlogsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func UpdateSpecResourcesBlogsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func DeleteSpecResourcesBlogsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}
