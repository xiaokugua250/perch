package sysadmin

import (
	"context"
	_ "fmt"
	"perch/pkg/sysinfo"
	"perch/web/metric"
	"perch/web/model"
	"strconv"
	"strings"
	"time"

	"net/http"
)

func SysBasicInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func SysMemInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func SysCpuInfoHandler(w http.ResponseWriter, req *http.Request) {
	metric.ProcessMetricFunc(w, req, nil)
}

func SysHostInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func SysDockerInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func SysDiskInfoHandler(w http.ResponseWriter, req *http.Request) {
	metric.ProcessMetricFunc(w, req, nil)
}

func SysNetInfoHandler(w http.ResponseWriter, req *http.Request) {
	metric.ProcessMetricFunc(w, req, nil)
}

func SysProcessInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func SysLoadInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}
