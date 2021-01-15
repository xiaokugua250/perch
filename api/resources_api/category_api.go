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

func GetResourcesCategorysHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceCategory []resource.ResourceCategory

			err          error
		)
		response.Kind = "Categorys"

		if err = database.MysqlDb.Find(&resourceCategory).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}
		if err = database.MysqlDb.Model(resource.ResourceCategory{}).Count(&response.Total).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = resourceCategory
		response.Message = " get resources Categorys successfully !!!"
		return nil
	})
}

func SpecGetResourcesCategorysHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceCategory resource.ResourceCategory
			DocID        string
			err          error
		)
		DocID = mux.Vars(r)["id"]

		response.Kind = "Categorys"

		if err = database.MysqlDb.Where("id=?", DocID).First(&resourceCategory).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}

		response.Total = 1
		response.Code = http.StatusOK
		response.Spec = resourceCategory
		response.Message = " get resources Categorys successfully !!!"
		return nil
	})
}

func CreateResourcesCategorysHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceCategory []resource.ResourceCategory
			err          error
		)
		response.Kind ="Categorys"

		if err = database.MysqlDb.Find(&resourceCategory).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}
		if err = database.MysqlDb.Model(resource.ResourceCategory{}).Count(&response.Total).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = resourceCategory
		response.Message = " create resources Categorys successfully !!!"
		return nil
	})
}

func UpdateSpecResourcesCategorysHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceCategory []resource.ResourceCategory
			DocID        string
			err          error
		)
		DocID = mux.Vars(r)["id"]
		response.Kind ="Categorys"

		if err = database.MysqlDb.Where("id=?", DocID).Updates(&resource.ResourceCategory{}).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = resourceCategory
		response.Message = " update resources Categorys successfully !!!"
		return nil
	})
}

func DeleteSpecResourcesCategorysHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			resourceCategory []resource.ResourceCategory
			DocID        string
			err          error
		)
		DocID = mux.Vars(r)["id"]
		if err = database.MysqlDb.Where("id=?", DocID).Delete(&resource.ResourceCategory{}).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = err.Error()
			return err
		}

		response.Total = 1
		response.Code = http.StatusOK
		response.Spec = resourceCategory
		response.Message = " delete resources Categorys successfully !!!"
		return nil
	})
}
