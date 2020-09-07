package sysadmin

import (
	"context"
	_ "fmt"
	"perch/pkg/sysinfo"
	"perch/web/metric"
	"perch/web/model"

	"net/http"
)

func SysMemInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			sysMemInfo sysinfo.SysMemInformation
			err         error
		)



		response.Kind = "sysinfo memory"

		sysMemInfo ,err=sysinfo.SysMemInfo()
		if err!= nil{
			response.Code=http.StatusInternalServerError
			response.Message=err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysMemInfo
		response.Message = " sys mem info"
		return nil
	})
}

func SysCpuInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			sysMemInfo sysinfo.SysMemInformation
			err         error
		)



		response.Kind = "sysinfo cpu"

		sysMemInfo ,err=sysinfo.SysMemInfo()
		if err!= nil{
			response.Code=http.StatusInternalServerError
			response.Message=err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysMemInfo
		response.Message = " sys mem info"
		return nil
	})
}

func SysHostInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			sysHostInfo sysinfo.HostAdvanceInfo
			err         error
		)



		response.Kind = "sysinfo memory"

		sysHostInfo ,err=sysinfo.SysHostAdvancedInfo()
		if err!= nil{
			response.Code=http.StatusInternalServerError
			response.Message=err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysHostInfo
		response.Message = " sys host info"
		return nil
	})
}

func SysDockerInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			sysMemInfo sysinfo.SysMemInformation
			err         error
		)



		response.Kind = "sysinfo memory"

		sysMemInfo ,err=sysinfo.SysMemInfo()
		if err!= nil{
			response.Code=http.StatusInternalServerError
			response.Message=err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysMemInfo
		response.Message = " sys mem info"
		return nil
	})
}

func SysDiskInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			sysMemInfo sysinfo.SysMemInformation
			err         error
		)



		response.Kind = "sysinfo memory"

		sysMemInfo ,err=sysinfo.SysMemInfo()
		if err!= nil{
			response.Code=http.StatusInternalServerError
			response.Message=err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysMemInfo
		response.Message = " sys mem info"
		return nil
	})
}

func SysNetInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			sysMemInfo sysinfo.SysMemInformation
			err         error
		)



		response.Kind = "sysinfo memory"

		sysMemInfo ,err=sysinfo.SysMemInfo()
		if err!= nil{
			response.Code=http.StatusInternalServerError
			response.Message=err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysMemInfo
		response.Message = " sys mem info"
		return nil
	})
}

func SysProcessInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			sysMemInfo sysinfo.SysMemInformation
			err         error
		)



		response.Kind = "sysinfo memory"

		sysMemInfo ,err=sysinfo.SysMemInfo()
		if err!= nil{
			response.Code=http.StatusInternalServerError
			response.Message=err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysMemInfo
		response.Message = " sys mem info"
		return nil
	})
}


func SysLoadInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			sysMemInfo sysinfo.SysMemInformation
			err         error
		)



		response.Kind = "sysinfo memory"

		sysMemInfo ,err=sysinfo.SysMemInfo()
		if err!= nil{
			response.Code=http.StatusInternalServerError
			response.Message=err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysMemInfo
		response.Message = " sys mem info"
		return nil
	})
}
