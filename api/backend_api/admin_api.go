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
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
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
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			user        rbac.AuthUser
			currentUser rbac.AuthUser
			err         error
		)
		response.Kind = "user token"
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
		token, err := auth.GenJwtToken(currentUser)
		if err != nil {
			response.Message = err.Error()
			response.Code = http.StatusInternalServerError
			return err
		}
		response.Code = http.StatusOK
		response.Spec = map[string]string{"token": token}
		response.Message = " user login successfully !!!"
		return nil
	})
}

//todo
func PlatLogoutHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		return nil
	})
}

func PlatUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			token string
		)
		query, err := url.ParseQuery(r.URL.RawQuery)
		if err != nil {
			response.Message = err.Error()
			return nil
		}
		token = query.Get("token")
		if token == "" {
			response.Message = "token is empty"
			response.Code = http.StatusBadRequest
			return nil
		}
		response.Code = http.StatusOK

		result := make(map[string]interface{})
		result["roles"] = []string{"admin"}
		result["introduction"] = "i am a super administrator ..."
		result["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
		result["name"] = "Super Admin"
		response.Spec = result
		response.Message = " Get User Info Successfully !!!"
		response.Kind = "user info"
		return nil
	})
}

func PlatAdminHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {

		return nil
	})

}
