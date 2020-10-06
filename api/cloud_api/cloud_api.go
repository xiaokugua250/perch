package sysadmin

import (
	"context"
	_ "fmt"
	"net/http"
	"perch/pkg/sysinfo"
	"perch/web/metric"
	"perch/web/model"
)

func CloudResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			sysMemInfo sysinfo.SysMemInformation
		//	err        error
		)

		response.Code = http.StatusOK
		response.Spec = sysMemInfo
		response.Message = " sys mem info"
		return nil
	})
}
