package api

import (
	"context"
	"net/http"
	"perch/internal/version"
	"perch/web/metric"
	"perch/web/model"
)

func ServiceHealthHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (

		//	err        error
		)

		response.Code = http.StatusOK
		response.Spec = map[string]string{"health": "ok"}
		response.Message = "ok"
		return nil
	})
}

func ServiceVersionandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (

		//	err        error
		)

		response.Code = http.StatusOK
		response.Spec = map[string]string{"version": version.Version, "compile": version.Compile, "author": version.Author, "email": version.Email}
		response.Message = "ok"
		return nil
	})
}
