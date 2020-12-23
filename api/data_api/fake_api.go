/**
 API
 假数据生成  https://godoc.org/github.com/icrowley/fake
 ref
*/
package data_api

import (
	"context"
	_ "fmt"
	"github.com/gorilla/mux"
	"net/http"
	database "perch/database/mysql"
	"perch/web/metric"
	"perch/web/model"
""github.com/icrowley/fake""
)

func GetFakeUsersHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs []resource.ResourceBlogs
			err          error
		)
		response.Kind = "docs"

		if err = database.MySQL_DB.Find(&resourceDocs).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}
		if err = database.MySQL_DB.Model(resource.ResourceBlogs{}).Count(&response.Total).Error; err != nil {
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

func GetFakeEmailsHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs resource.ResourceBlogs
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


func GetFakeCrediCardHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs resource.ResourceBlogs
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


func GetFakeIPHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs resource.ResourceBlogs
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


func GetFakeLocHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs resource.ResourceBlogs
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
func GetFakeTimesHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceDocs resource.ResourceBlogs
			DocID        string
			err          error
		)
		DocID = mux.Vars(r)["id"]

		response.Kind = "docs"

		fake.
		response.Total = 1
		response.Code = http.StatusOK
		response.Spec = resourceDocs
		response.Message = " get resources docs successfully !!!"
		return nil
	})
}
