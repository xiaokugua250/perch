package auth_api

import (
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	database "perch/database/mysql"
	casbin "perch/pkg/auth/casbin"
	"perch/web/metric"
	"perch/web/model"

	"strconv"
)

func CasbinAuthSpcGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			userCasbinsRules casbin.CasbinRule
			id               int
			err              error
		)
		response.Kind = "casbin auths"
		id, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		if err = database.MysqlDb.Where("id=?", id).First(userCasbinsRules).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = userCasbinsRules
		response.Message = " get auth user casbin rules successfully"
		return nil
	})
}

func CasbinAuthGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			userCasbinsRules []casbin.CasbinRule

			err error
		)
		response.Kind = "casbin auths"

		if err = database.MysqlDb.Find(&userCasbinsRules).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		if err = database.MysqlDb.Model(&casbin.CasbinRule{}).Count(&response.Total).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		response.Code = http.StatusOK
		response.Spec = userCasbinsRules
		response.Message = " get auth users casbin rules successfully"
		return nil
	})
}

func CasbinAuthCreateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			userCasbinsRule casbin.CasbinRule

			err error
		)
		response.Kind = "casbin auths"

		if err = json.NewDecoder(r.Body).Decode(&userCasbinsRule); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()
			return err
		}
		if err = database.MysqlDb.Create(userCasbinsRule).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = userCasbinsRule
		response.Message = " get auth users successfully"
		return nil
	})
}

func CasbinAuthPatchHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			userCasbinsRule casbin.CasbinRule
			id              int
			err             error
		)
		response.Kind = "casbin auths"
		id, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		if err = json.NewDecoder(r.Body).Decode(&userCasbinsRule); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()
			return err
		}
		if err = database.MysqlDb.Where("id=?", id).Model(casbin.CasbinRule{}).Updates(userCasbinsRule).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = userCasbinsRule
		response.Message = "update auth users casbin successfully"
		return nil
	})
}

func CasbinAuthDeleteHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			userCasbinsRule casbin.CasbinRule
			id              int
			err             error
		)
		response.Kind = "casbin auths"
		id, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		if err = database.MysqlDb.Where("id=?", id).Delete(&userCasbinsRule).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		response.Code = http.StatusOK
		response.Spec = userCasbinsRule
		response.Message = " delete user casbin role successfully"
		return nil
	})
}
