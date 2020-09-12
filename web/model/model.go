package model
import (
	"github.com/dgrijalva/jwt-go"

)
type BaseReponse struct {
	Message string `json:"message"`
	Kind    string `json:"kind"`
	Code    int    `json:"code"`
	Total   int64    `json:"total"`
}
type ResultReponse struct {
	BaseReponse
	Spec interface{} `json:"spec"`
}


type PEXToken struct {
	UserID             int      `json:"id"`
	UserName           string   `json:"user_name"` // 用户名
	UserRoles          []string `json:"user_roles"`
	Status             int      `json:"status"` // 用户状态
	jwt.StandardClaims          // 标准JWT包含的Token内容
}
