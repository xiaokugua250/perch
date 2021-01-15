package admin

import (
	"context"
	"encoding/json"
	_ "fmt"
	"net/http"
	"net/url"
	database "perch/database/mysql"
	"perch/web/auth"
	"perch/web/metric"
	"perch/web/model"
	rbac "perch/web/model/rbac"
)

func PlatLoginHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, func(ctx context.Context, bean interface{}, response *model.ResultReponse) error {
		var (
			user        rbac.AuthUser
			currentUser rbac.AuthUser
			err         error
		)
		response.Kind = "user"
		if err = json.NewDecoder(r.Body).Decode(&user); err != nil {
			response.Code = http.StatusBadRequest
			response.Message = "user login failed !!!"
			response.Spec = currentUser
			return err
		}

		if err = database.MysqlDb.Where("username=?", user.UserName).First(&currentUser).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = user
			return err
		}

		if user.UserPasswd != currentUser.UserPasswd {
			response.Code = http.StatusBadRequest
			response.Message = "user login failed !!!"
			response.Spec = currentUser
			return err
			//response.Message =" user login successfully !!!"
		}

		response.Code = http.StatusOK
		response.Spec = map[string]string{"token": "admin-token"}
		response.Message = " user login successfully !!!"
		return nil
	})
}

func PlatLoginGenTokenHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

//todo
func PlatLogoutHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func PlatUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)
}

func PlatAdminHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil)

}
