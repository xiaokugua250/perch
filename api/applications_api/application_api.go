package applications_api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	database "perch/database/mysql"

	"perch/web/metric"
	"perch/web/model"
	"perch/web/model/applications"

	"strconv"
)

/**

创建软件应用

*/
func ApplicationsCreateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			application    applications.Applications
			isAlreadyExist int64

			err error
		)
		response.Kind = "application"
		if err = json.NewDecoder(r.Body).Decode(&application); err != nil {
			response.Code = http.StatusBadRequest
			return err
		}

		if err = database.MysqlDb.Model(applications.Applications{}).Where("name=?", application.Name).Count(&isAlreadyExist).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = application
			return err
		}
		if isAlreadyExist > 0 {
			return errors.New(fmt.Sprintf(" application with name %s already exits..."))
		}
		if err = database.MysqlDb.Create(&application).Error; err != nil {
			return err
		}

		response.Spec = application

		response.Code = http.StatusOK
		response.Total = 1
		response.Message = "create applications successfully !!!"
		return nil
	})
}

/**
查询应用

*/
func ApplicationsGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			apps []applications.Applications

			err error
		)
		response.Kind = "applications"

		if err = database.MysqlDb.Find(&apps).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()

			return err
		}
		if err = database.MysqlDb.Model(applications.Applications{}).Count(&response.Total).Error; err != nil {
			return err
		}

		response.Spec = apps

		response.Code = http.StatusOK

		response.Message = "get all applications successfully !!!"
		return nil
	})
}

/**
查询特定应用
*/
func ApplicationsSpecGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			application applications.Applications
			id          int

			err error
		)
		response.Kind = "application"
		id, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return err
		}

		if err = database.MysqlDb.Where("id=?", id).First(&application).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Spec = application
		response.Code = http.StatusOK
		response.Total = 1
		response.Message = "get spec application successfully !!!"
		return nil
	})
}

/**

更新特定应用

*/
func ApplicationsSpecUpdateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			application applications.Applications
			id          int

			err error
		)
		response.Kind = "application"
		id, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return err
		}
		if err = json.NewDecoder(r.Body).Decode(&application); err != nil {
			response.Code = http.StatusBadRequest
			return err
		}
		application.ID = id

		if err = database.MysqlDb.Save(&application).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Total = 1
		response.Message = " application update successfully !!!"
		return nil
	})
}

/**
删除特定应用

*/
func ApplicationsSpecDeleteHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			application applications.Applications
			id          int

			err error
		)
		response.Kind = "application"
		id, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			return err
		}

		if err = database.MysqlDb.Where("id=?", id).Delete(application).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Total = 1
		response.Message = "application delete successfully !!!"
		return nil
	})
}
