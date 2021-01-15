package messages

import (
	"context"
	_ "fmt"
	"io"
	"net/http"
	"os"
	"perch/pkg/cluster/k8s"
	"perch/pkg/sysinfo"
	"perch/web/metric"
	"perch/web/model"
	"time"
)

func CloudResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

/**
根据资源文件进行k8s 集群资源处理
*/
func CloudResourceFileHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}
