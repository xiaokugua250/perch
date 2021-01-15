package admin

import (
	"context"
	"encoding/json"
	_ "fmt"
	"net/http"
	"net/url"
	database "perch/database/mysql"
	"perch/web/auth"
	"perch/web/metric"
	"perch/web/model"
	rbac "perch/web/model/rbac"
)

func PlatLoginHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func PlatLoginGenTokenHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

//todo
func PlatLogoutHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func PlatUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func PlatAdminHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)

}
