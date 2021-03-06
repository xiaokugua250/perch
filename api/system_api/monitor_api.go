package sysadmin

import (
	"context"
	_ "fmt"
	"perch/pkg/monitor/system"

	"perch/web/metric"
	"perch/web/model"
	"strconv"
	"strings"
	"time"

	"net/http"
)

func SysBasicInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			//sysBasicInfo = make(map[string]interface{})
			sysBasicInfo system.HostAdvancedInfo
			err          error
		)
		response.Kind = "sysinfo basic"
		sysBasicInfo, err = system.SysHostAdvancedInfo()
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysBasicInfo
		response.Message = "sys basic info"
		return nil
	})
}

func SysMemInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			sysMemInfo system.SysMemInformation
			err        error
		)

		response.Kind = "sysinfo memory"

		sysMemInfo, err = system.SysMemInfo()
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysMemInfo
		response.Message = " sys mem info"
		return nil
	})
}

func SysCpuInfoHandler(w http.ResponseWriter, req *http.Request) {
	metric.ProcessMetricFunc(w, req, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			sysCpuInfo system.CpuAdvancedInfo
			logical    bool
			percpu     bool
			interval   time.Duration
			err        error
		)
		logicalStr := req.URL.Query().Get("logical")
		if logicalStr == "" {
			logical = true
		} else {
			logical, err = strconv.ParseBool(logicalStr)
			if err != nil {
				response.Code = http.StatusBadRequest
				response.Message = err.Error()
				return err
			}
		}

		percpuStr := req.URL.Query().Get("percpu")
		if percpuStr == "" {
			percpu = true
		} else {
			percpu, err = strconv.ParseBool(percpuStr)
			if err != nil {
				response.Code = http.StatusBadRequest
				response.Message = err.Error()
				return err
			}
		}

		intervalStr := req.URL.Query().Get("interval")
		if intervalStr == "" {
			interval = 1 * time.Second
		} else {
			interval, err = time.ParseDuration(intervalStr)
			if err != nil {
				response.Code = http.StatusBadRequest
				response.Message = err.Error()
				return err
			}
		}

		response.Kind = "sysinfo cpu"

		sysCpuInfo, err = system.SysAdvancedCpuInfo(logical, percpu, interval)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysCpuInfo
		response.Total = 1
		response.Message = " sys mem info"
		return nil
	})
}

func SysHostInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			sysHostInfo system.HostAdvancedInfo
			err         error
		)

		response.Kind = "sysinfo memory"

		sysHostInfo, err = system.SysHostAdvancedInfo()
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysHostInfo
		response.Message = " sys host info"
		return nil
	})
}

func SysDockerInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			sysDockerInfo system.DockerAdvancedInfo
			err           error
		)

		response.Kind = "sysinfo docker"

		sysDockerInfo, err = system.SysAdvancedDockerInfo()
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysDockerInfo
		response.Total = 1
		response.Message = " sys docker info"
		return nil
	})
}

func SysDiskInfoHandler(w http.ResponseWriter, req *http.Request) {
	metric.ProcessMetricFunc(w, req, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			diskSerialName string
			diskLableName  string
			partions       bool
			path           string
			iocounters     []string
			sysDiskInfo    system.DiskAdvacedInfo

			err error
		)

		diskSerialName = req.URL.Query().Get("diskSerialName")
		diskLableName = req.URL.Query().Get("diskLableName")
		partionsStr := req.URL.Query().Get("partions")
		if partionsStr == "" {
			partions = true
		} else {
			partions, err = strconv.ParseBool(partionsStr)
			if err != nil {
				response.Message = err.Error()
				response.Code = http.StatusBadRequest
				return err
			}

		}

		path = req.URL.Query().Get("path")
		iocountersStr := req.URL.Query().Get("iocounters")
		iocounters = strings.Split(iocountersStr, ",")
		response.Kind = "sysinfo disk"

		sysDiskInfo, err = system.SysAdvancedDiskInfo(diskSerialName, diskLableName, partions, path, iocounters...)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Total = 1
		response.Code = http.StatusOK
		response.Spec = sysDiskInfo
		response.Message = " sys disk info"
		return nil
	})
}

func SysNetInfoHandler(w http.ResponseWriter, req *http.Request) {
	metric.ProcessMetricFunc(w, req, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			sysNetInfo system.NetAdvancedInfo
			percpu     bool
			err        error
		)

		response.Kind = "sysinfo net"
		percpuStr := req.URL.Query().Get("percpu")
		if percpuStr == "" {
			percpu = true
		} else {
			percpu, err = strconv.ParseBool(percpuStr)
			if err != nil {
				response.Code = http.StatusBadRequest
				response.Message = err.Error()
				return err
			}
		}
		sysNetInfo, err = system.SysAdvancedNetInfo(percpu)
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysNetInfo
		response.Message = " sys net info"
		return nil
	})
}

func SysProcessInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			sysProcessInfo system.ProcessAdvancedInfo
			err            error
		)

		response.Kind = "sysinfo process"

		sysProcessInfo, err = system.SysAdvancedProcessInfo()
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysProcessInfo
		response.Message = " sys process info"
		return nil
	})
}

func SysLoadInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			sysLoadInfo system.LoadAdvancedInfo
			err         error
		)

		response.Kind = "sysinfo memory"

		sysLoadInfo, err = system.SysAdvancedLoadInfo()
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = sysLoadInfo
		response.Message = " sys load info"
		return nil
	})
}
