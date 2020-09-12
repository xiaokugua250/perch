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

func PlatAuthRolesGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userRoles      []rbac.AuthRBACRoles

			err         error
		)
		response.Kind = "auth rbac roles"

		if err = database.MySQL_DB.Find(&userRoles).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		if err = database.MySQL_DB.Model(&rbac.AuthRBACRoles{}).Count(&response.Total).Error;err!= nil{
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		response.Code = http.StatusOK
		response.Spec =	userRoles
		response.Message = " get auth rbac roles successfully"
		return nil
	})
}

func PlatSpecAuthRoleGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userRoles       rbac.AuthRBACRoles
			roleID  int
			err         error
		)


		response.Kind = "auth user role"
		roleID,err = strconv.Atoi(mux.Vars(r)["id"])
		if err!= nil{
			response.Code= http.StatusBadRequest
			response.Message= err.Error()
			return err
		}

		if err = database.MySQL_DB.Where("id=?",roleID).First(userRoles).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = userRoles
			return err
		}

		response.Total=1
		response.Code = http.StatusOK
		response.Spec = userRoles
		response.Message = " get spec auth users roles successfully !!!"
		return nil
	})
}

func PlatAuthRoleUpdateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userRoles       rbac.AuthRBACRoles
			roleID  int
			err         error
		)
		response.Kind = "auth user role"
		if err = json.NewDecoder(r.Body).Decode(&userRoles); err != nil {
			response.Code = http.StatusBadRequest
			response.Message =err.Error()
			return err
		}
		response.Kind = "auth user role"
		roleID,err = strconv.Atoi(mux.Vars(r)["id"])
		if err!= nil{
			response.Code= http.StatusBadRequest
			response.Message= err.Error()
			return err
		}

		if err = database.MySQL_DB.Where("id=?",roleID).Updates(userRoles).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = userRoles
			return err
		}


		response.Code = http.StatusOK
		response.Spec = userRoles
		response.Message = " update  auth users roles successfully !!!"
		return nil
	})
}
func PlatAuthRoleDeleteHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userRoles       rbac.AuthRBACRoles
			roleID  int
			err         error
		)
		response.Kind = "auth user"

		roleID ,err = strconv.Atoi(mux.Vars(r)["id"])
		if err!= nil{
			response.Code= http.StatusBadRequest
			response.Message= err.Error()
			return err
		}

		if err = database.MySQL_DB.Where("id=?", 	roleID ).Delete(&rbac.AuthRBACRoles{}).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = userRoles
			return err
		}

		response.Total=1
		response.Code = http.StatusOK
		//response.Spec = "user with uid "+userID+" delete "
		response.Message = " delete auth users roles successfully!!!"
		return nil
	})
}
func PlatAuthRoleCreateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userRoles      rbac.AuthRBACRoles
			err         error
		)
		response.Kind = "auth user roles"
		if err = json.NewDecoder(r.Body).Decode(&userRoles); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = "user create failed !!!"
			response.Spec = userRoles
			return err
		}

		if err = database.MySQL_DB.Create(&userRoles).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = userRoles
			return err
		}
		response.Total=1

		response.Code = http.StatusOK
		response.Spec = userRoles
		response.Message = "  create auth users roles successfully"
		return nil
	})
}


//todo 获取特定用户的权限
func PlatAuthSpecRolePermissionsGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userRoles      rbac.AuthRBACRoles
			err         error
		)
		response.Kind = "auth user roles"
		if err = json.NewDecoder(r.Body).Decode(&userRoles); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = "user create failed !!!"
			response.Spec = userRoles
			return err
		}

		if err = database.MySQL_DB.Create(&userRoles).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = userRoles
			return err
		}
		response.Total=1

		response.Code = http.StatusOK
		response.Spec = userRoles
		response.Message = "  create auth users roles successfully"
		return nil
	})
}


func PlatAuthPermissionsGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userPermissions       []rbac.AuthRBACPermissions

			err         error
		)
		response.Kind = "auth users"

		if err = database.MySQL_DB.Find(&	userPermissions).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		if err = database.MySQL_DB.Model(&rbac.AuthRBACPermissions{}).Count(&response.Total).Error;err!= nil{
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		response.Total=1
		response.Code = http.StatusOK
		response.Spec =	userPermissions
		response.Message = " get auth users successfully"
		return nil
	})
}


func PlatSpecAuthPermissionGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userPermissions     rbac.AuthRBACPermissions
			permissionID  int
			err         error
		)
		response.Kind = "auth user"


		permissionID ,err = strconv.Atoi(mux.Vars(r)["id"])
		if err!= nil{
			response.Code= http.StatusBadRequest
			response.Message= err.Error()
			return err
		}

		if err = database.MySQL_DB.Where("id=?",	permissionID ).First(	&userPermissions).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = 	userPermissions
			return err
		}


		response.Code = http.StatusOK
		response.Spec = 	userPermissions
		response.Total=1
		response.Message = " get spec auth users permissions  successfully !!!"
		return nil
	})
}
func PlatAuthPermissionUpdateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userPermissions     rbac.AuthRBACPermissions
			permissionID  int
			err         error
		)
		response.Kind = "auth user"
		if err = json.NewDecoder(r.Body).Decode(&userPermissions); err != nil {
			response.Code = http.StatusBadRequest
			response.Message =err.Error()
			return err
		}

		permissionID ,err = strconv.Atoi(mux.Vars(r)["id"])
		if err!= nil{
			response.Code= http.StatusBadRequest
			response.Message= err.Error()
			return err
		}

		if err = database.MySQL_DB.Where("id=?",	permissionID ).Updates(	&userPermissions).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = 	userPermissions
			return err
		}

		response.Total=1
		response.Code = http.StatusOK
		response.Spec = 	userPermissions
		response.Message = " update  auth users permissions  successfully !!!"
		return nil
	})
}


func PlatAuthPermissionDeleteHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userPermissions   rbac.AuthRBACPermissions
			permissionID   int
			err         error
		)
		response.Kind = "auth user"

		permissionID ,err = strconv.Atoi(mux.Vars(r)["id"])
		if err!= nil{
			response.Code= http.StatusBadRequest
			response.Message= err.Error()
			return err
		}

		if err = database.MySQL_DB.Where("id=?", 	permissionID ).Delete(&rbac.AuthRBACPermissions{}).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = 	userPermissions
			return err
		}

		response.Total=1
		response.Code = http.StatusOK
		//response.Spec = "user with uid "+userID+" delete "
		response.Message = " delete auth users  permissions successfully!!!"
		return nil
	})
}
func PlatAuthPermissionCreateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			userPermissions    rbac.AuthRBACPermissions

			err         error
		)
		response.Kind = "auth user"
		if err = json.NewDecoder(r.Body).Decode(&	userPermissions); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = "user create failed !!!"

			return err
		}

		if err = database.MySQL_DB.Create(&	userPermissions).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = 	userPermissions
			return err
		}

		response.Total=1
		response.Code = http.StatusOK
		response.Spec = 	userPermissions
		response.Message = "  create auth users permissions successfully"
		return nil
	})
}
