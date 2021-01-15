package api

import (
	"context"
	"net/http"
	"perch/web/metric"
	"perch/web/model"
)

func ServiceHealthHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, respone *model.ResultResponse) error {
		return nil
	})
}

func ServiceVersionandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, respone *model.ResultResponse) error {
		return nil
	})
}
