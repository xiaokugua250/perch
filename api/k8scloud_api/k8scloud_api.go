package k8scloud

import (
	"context"
	_ "fmt"
	"io"
	"net/http"
	"os"

	"perch/pkg/sysinfo"
	"perch/web/metric"
	"perch/web/model"
	"time"
)

func CloudResoucesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			sysMemInfo monitor.SysMemInformation
			//	err        error
		)

		response.Code = http.StatusOK
		response.Spec = sysMemInfo
		response.Message = " sys mem info"
		return nil
	})
}

/**
根据资源文件进行k8s 集群资源处理
*/
func CloudResourceFileHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			//	yamlFile string
			err error
		)
		file, fileHeader, err := r.FormFile("yamlfile")
		if err != nil {
			return err
		}
		defer file.Close()
		targetFileName := fileHeader.Filename + time.Now().String()
		targetFile, err := os.OpenFile("/pathToStorageFile/"+targetFileName, os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		io.Copy(targetFile, file)
		//yamlFile = "/pathToStorageFile/" + targetFileName
		//	err = k8scloud.K8SClientSet.K8SConstructorFileValidate(yamlFile)
		response.Code = http.StatusOK
		response.Spec = "k8s resource created by file successfully!!"
		response.Message = "k8s resource created by file successfully!!"
		return nil
	})
}
