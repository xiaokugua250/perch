/**
文本资源 API
*/
package resources_api

import (
	"context"
	_ "fmt"
	"github.com/gorilla/mux"
	"net/http"
	database "perch/database/mysql"
	"perch/web/metric"
	"perch/web/model"
	resource "perch/web/model/resources"
)

func GetResourcesDocsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs []resource.ResourceDocs
			err          error
		)
		response.Kind = "docs"

		if err = database.MySQL_DB.Find(&resourceDocs).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}
		if err = database.MySQL_DB.Model(resource.ResourceDocs{}).Count(&response.Total).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = resourceDocs
		response.Message = " get resources docs successfully !!!"
		return nil
	})
}

func SpecGetResourcesDocsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs resource.ResourceDocs
			DocID        string
			err          error
		)
		DocID = mux.Vars(r)["id"]

		response.Kind = "docs"

		if err = database.MySQL_DB.Where("id=?", DocID).First(&resourceDocs).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}

		response.Total = 1
		response.Code = http.StatusOK
		response.Spec = resourceDocs
		response.Message = " get resources docs successfully !!!"
		return nil
	})
}

func CreateResourcesDocsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs []resource.ResourceDocs
			err          error
		)
		response.Kind ="docs"

		if err = database.MySQL_DB.Find(&resourceDocs).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}
		if err = database.MySQL_DB.Model(resource.ResourceDocs{}).Count(&response.Total).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = resourceDocs
		response.Message = " create resources docs successfully !!!"
		return nil
	})
}

func UpdateSpecResourcesDocsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs []resource.ResourceDocs
			DocID        string
			err          error
		)
		DocID = mux.Vars(r)["id"]
		response.Kind ="docs"

		if err = database.MySQL_DB.Where("id=?", DocID).Updates(&resource.ResourceDocs{}).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = resourceDocs
		response.Message = " update resources docs successfully !!!"
		return nil
	})
}

func DeleteSpecResourcesDocsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs []resource.ResourceDocs
			DocID        string
			err          error
		)
		DocID = mux.Vars(r)["id"]
		if err = database.MySQL_DB.Where("id=?", DocID).Delete(&resource.ResourceDocs{}).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}

		response.Total = 1
		response.Code = http.StatusOK
		response.Spec = resourceDocs
		response.Message = " delete resources docs successfully !!!"
		return nil
	})
}
