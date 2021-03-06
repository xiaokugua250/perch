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

func GetResourcesBlogsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			resourceDocs []resource.ResourceBlogs
			err          error
		)
		response.Kind = "docs"

		if err = database.MysqlDb.Find(&resourceDocs).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}
		if err = database.MysqlDb.Model(resource.ResourceBlogs{}).Count(&response.Total).Error; err != nil {
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

func SpecGetResourcesBlogsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			resourceDocs resource.ResourceBlogs
			DocID        string
			err          error
		)
		DocID = mux.Vars(r)["id"]

		response.Kind = "docs"

		if err = database.MysqlDb.Where("id=?", DocID).First(&resourceDocs).Error; err != nil {
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

func CreateResourcesBlogsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			resourceDocs []resource.ResourceBlogs
			err          error
		)
		response.Kind = "docs"

		if err = database.MysqlDb.Find(&resourceDocs).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}
		if err = database.MysqlDb.Model(resource.ResourceBlogs{}).Count(&response.Total).Error; err != nil {
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

func UpdateSpecResourcesBlogsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			resourceDocs []resource.ResourceBlogs
			DocID        string
			err          error
		)
		DocID = mux.Vars(r)["id"]
		response.Kind = "docs"

		if err = database.MysqlDb.Where("id=?", DocID).Updates(&resource.ResourceBlogs{}).Error; err != nil {
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

func DeleteSpecResourcesBlogsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			resourceDocs []resource.ResourceBlogs
			DocID        string
			err          error
		)
		DocID = mux.Vars(r)["id"]
		if err = database.MysqlDb.Where("id=?", DocID).Delete(&resource.ResourceBlogs{}).Error; err != nil {
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
