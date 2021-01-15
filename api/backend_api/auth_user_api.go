package admin

import (
	"context"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"net/http"
	database "perch/database/mysql"
	"perch/web/metric"
	"perch/web/model"
	rbac "perch/web/model/rbac"

	"strconv"
)

func PlatAuthUsersGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

//todo 需要获取到用户角色，权限等信息
func PlatSpecAuthUserGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}
func PlatAuthUserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}
func PlatAuthUserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}
func PlatAuthUserCreateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}
