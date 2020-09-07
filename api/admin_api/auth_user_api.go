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
	"strconv"
)

func PlatAuthUsersGetHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			user       []model.AuthUser

			err         error
		)
		response.Kind = "auth users"

		if err = database.MySQL_DB.Find(&user).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}

		if err = database.MySQL_DB.Model(&model.AuthUser{}).Count(&response.Total).Error;err!= nil{
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			return err
		}
		response.Code = http.StatusOK
		response.Spec =user
		response.Message = " get auth users successfully"
		return nil
	})
}

func PlatAuthUserUpdateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			user        model.AuthUser
			userID  int
			err         error
		)
		response.Kind = "auth user"
		if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.Code = http.StatusBadRequest
			response.Message =err.Error()
			return err
		}
		response.Kind = "auth user"
		userID,err = strconv.Atoi(mux.Vars(r)["id"])
		if err!= nil{
			response.Code= http.StatusBadRequest
			response.Message= err.Error()
			return err
		}

		if err = database.MySQL_DB.Where("id=?",userID).Updates(user).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = user
			return err
		}


		response.Code = http.StatusOK
		response.Spec = user
		response.Message = " update  auth users successfully !!!"
		return nil
	})
}
func PlatAuthUserDeleteHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			user        model.AuthUser
			userID  int
			err         error
		)
		response.Kind = "auth user"

		userID,err = strconv.Atoi(mux.Vars(r)["id"])
		if err!= nil{
			response.Code= http.StatusBadRequest
			response.Message= err.Error()
			return err
		}

		if err = database.MySQL_DB.Where("id=?", userID).Delete(&model.AuthUser{}).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = user
			return err
		}


		response.Code = http.StatusOK
		//response.Spec = "user with uid "+userID+" delete "
		response.Message = " delete auth users successfully!!!"
		return nil
	})
}
func PlatAuthUserCreateHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			user        model.AuthUser
			currentUser model.AuthUser
			err         error
		)
		response.Kind = "auth user"
		if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = "user create failed !!!"
			response.Spec = currentUser
			return err
		}

		if err = database.MySQL_DB.Create(&user).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = user
			return err
		}


		response.Code = http.StatusOK
		response.Spec = user
		response.Message = "  create auth users successfully"
		return nil
	})
}
