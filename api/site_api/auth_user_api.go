package admin

import (
	"context"
	"encoding/json"
	_ "fmt"
	"github.com/gorilla/mux"
	"net/http"
	database "perch/database/mysql"
	"perch/web/metric"
	"perch/web/model"
	rbac "perch/web/model/rbac"

	"strconv"
)

func PlatAuthUsersGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			user []rbac.AuthUser

			err error
		)
		response.Kind = "auth users"

		if err = database.MysqlDb.Find(&user).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		if err = database.MysqlDb.Model(&rbac.AuthUser{}).Count(&response.Total).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		response.Code = http.StatusOK
		response.Spec = user
		response.Message = " get auth users successfully"
		return nil
	})
}

//todo 需要获取到用户角色，权限等信息
func PlatSpecAuthUserGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			user      rbac.AuthUser
			userRoles []rbac.AuthRBACRoles
			userID    int
			err       error
		)
		response.Kind = "auth user"

		userID, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()
			return err
		}

		if err = database.MysqlDb.Where("id=?", userID).First(&user).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = user
			return err
		}
		//	subQuery := database.MySQL_DB.Table("auth_rbac_user_roles").Select("role_id").Where("user_id=?",userID)
		subQuery := database.MysqlDb.Model(rbac.AuthRBACUserRoles{}).Select("role_id").Where("user_id=?", userID)

		if err = database.MysqlDb.Where("id in (?)", subQuery).Find(&userRoles).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = nil
			return err
		}
		user.UserRoles = userRoles
		response.Code = http.StatusOK
		response.Spec = user
		response.Total = 1
		response.Message = " get spec auth users successfully !!!"
		return nil
	})
}
func PlatAuthUserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			user   rbac.AuthUser
			userID int
			err    error
		)
		response.Kind = "auth user"
		if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()
			return err
		}
		response.Kind = "auth user"
		userID, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()
			return err
		}

		if err = database.MysqlDb.Where("id=?", userID).Updates(user).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = user
			return err
		}

		response.Total = 1
		response.Code = http.StatusOK
		response.Spec = user
		response.Message = " update  auth users successfully !!!"
		return nil
	})
}
func PlatAuthUserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			user   rbac.AuthUser
			userID int
			err    error
		)
		response.Kind = "auth user"

		userID, err = strconv.Atoi(mux.Vars(r)["id"])
		if err != nil {
			response.Code = http.StatusBadRequest
			response.Message = err.Error()
			return err
		}

		if err = database.MysqlDb.Where("id=?", userID).Delete(&rbac.AuthUser{}).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = user
			return err
		}

		response.Total = 1
		response.Code = http.StatusOK
		//response.Spec = "user with uid "+userID+" delete "
		response.Message = " delete auth users successfully!!!"
		return nil
	})
}
func PlatAuthUserCreateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			user        rbac.AuthUser
			currentUser rbac.AuthUser
			err         error
		)
		response.Kind = "auth user"
		if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = "user create failed !!!"
			response.Spec = currentUser
			return err
		}

		if err = database.MysqlDb.Create(&user).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = user
			return err
		}

		response.Total = 1
		response.Code = http.StatusOK
		response.Spec = user
		response.Message = "  create auth users successfully"
		return nil
	})
}
