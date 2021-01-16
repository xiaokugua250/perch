package model
import (
	"github.com/dgrijalva/jwt-go"

)
type BasicResponse struct {
	Message string `json:"message"`
	Kind    string `json:"kind"`
	Code    int    `json:"code"`
	Total   int64    `json:"total"`
}
type ResultResponse struct {
	BasicResponse
	SecretToken `json:"-"`
	Spec interface{} `json:"spec"`
}


type SecretToken struct {
	//UserID             int      `json:"id"`
	UserName           string   `json:"user_name"` // 用户名
	//UserRoles          []string `json:"user_roles"`
	UserUID int64 `json:"user_uid"`
	UserGID int64 `json:"user_gid"`
	Status             int64      `json:"status"` // 用户状态
	jwt.StandardClaims          // 标准JWT包含的Token内容
}
