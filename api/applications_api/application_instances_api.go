package applications_api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	database "perch/database/mysql"

	"perch/web/metric"
	"perch/web/model/applications"
	"strconv"

	"perch/web/model"
)

/**
生成应用实例
*/
func ApplicationsInstancesCreateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			instance applications.ApplicationInstances

			err error
		)
		response.Kind = "application_instances"
		if err = json.NewDecoder(r.Body).Decode(&instance); err != nil {
			response.Code = http.StatusBadRequest
			return err
		}



		if err = database.MysqlDb.Create(&instance).Error; err != nil {
			return err
		}

		response.Spec = instance

		response.Code = http.StatusOK
		response.Total = 1
		response.Message = "create application instance successfully !!!"
		return nil
	})
}


/**

查询应用实例

*/
func ApplicationsInstancesGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			instances  []applications.ApplicationInstances

			err         error
		)
		response.Kind = "application_instances"


		if err = database.MysqlDb.Find(&instances).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()

			return err
		}
		if err= database.MysqlDb.Model(&applications.ApplicationInstances{}).Count(&response.Total).Error;err!= nil{
			return err
		}


		response.Spec = instances

		response.Code = http.StatusOK

		response.Message = "get all application instance successfully !!!"
		return nil
	})
}



/**
获取特定应用实例
 */
func ApplicationsInstancesSpecGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			instance applications.ApplicationInstances
			id          int

			err error
		)
		response.Kind = "application_instances"
		id, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return err
		}

		if err = database.MysqlDb.Where("id=?", id).First(&instance).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Spec = instance
		response.Code = http.StatusOK
		response.Total = 1
		response.Message = "get spec application instance successfully !!!"
		return nil
	})
}

/**
更新应用实例

*/
func ApplicationsInstancesSpecUpdateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			instance applications.ApplicationInstances
			id          int

			err error
		)
		response.Kind = "application_instances"
		id, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return err
		}
		if err = json.NewDecoder(r.Body).Decode(&instance); err != nil {
			response.Code = http.StatusBadRequest
			return err
		}
		instance.ID = id

		if err = database.MysqlDb.Save(&instance).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Total = 1
		response.Message = " application instance update successfully !!!"
		return nil
	})
}


/**
删除应用实例

*/
func ApplicationsInstancesSpecDeleteHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			instance applications.ApplicationInstances
			id          int

			err error
		)
		response.Kind = "application_instances"
		id, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return err
		}

		if err = database.MysqlDb.Where("id=?", id).Delete(instance).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Total = 1
		response.Message = "application instance delete successfully !!!"
		return nil
	})
}


