package admin

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"perch/pkg/general/utils/secure"
	"time"

	_ "fmt"
	"net/http"
	database "perch/database/mysql"
	"perch/web/auth"
	"perch/web/metric"
	"perch/web/model"
	rbac "perch/web/model/rbac"
)

/**
用户登录，登录成功则生成用户token

*/
func AuthUserSignInHandler(w http.ResponseWriter, r *http.Request) {
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
			return err
		}

		if err = database.MysqlDb.Where("user_name=?", user.UserName).First(&currentUser).Error; err != nil {
			response.Code = http.StatusInternalServerError
			response.Message = err.Error()
			response.Spec = user
			return err
		}

		userPasswd := secure.GenerateSHA1Hash(secure.GenerateMd5Hash(user.UserPasswd + currentUser.UserSalt))
		if userPasswd != currentUser.UserPasswd {
			return errors.New(fmt.Sprintf("user %s login failed,please check your user_name or passwd...", user.UserName))
		}
		userToken, err := auth.GenJwtToken(currentUser)
		if err != nil {
			return err
		}
		if err= database.MysqlDb.Model(&rbac.AuthUser{}).Where("user_name=?",user.UserName).Update("last_login",time.Now().Unix()).Error;err!= nil{
			return err
		}
		response.Spec = userToken

		response.Code = http.StatusOK
		response.Total = 1
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

func AuthUserInfoHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{AuthPlugin: metric.AuthPlugin{AuthToken: true}}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {
		var (
			err error
			user rbac.AuthUser
		)
		if err= database.MysqlDb.Where("user_name=?",response.SecretToken.UserName).First(&user).Error;err!= nil{
			return err
		}

		response.Spec=user
		response.Total=1
		response.Code=http.StatusOK
		/*result := make(map[string]interface{})
		result["roles"] = []string{"admin"}
		result["introduction"] = "i am a super administrator ..."
		result["avatar"] = "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif"
		result["name"] = "Super Admin"
		response.Spec = result
		response.Message = " Get User Info Successfully !!!"*/
		response.Kind = "user info"
		return nil
	})
}

func PlatAdminHandler(w http.ResponseWriter, r *http.Request) {
	metric.ProcessMetricFunc(w, r, nil, &metric.MiddlewarePlugins{}, func(ctx context.Context, bean interface{}, response *model.ResultResponse) error {

		return nil
	})

}
