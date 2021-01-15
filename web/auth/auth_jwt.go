package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"perch/web/model"
	rbac "perch/web/model/rbac"
	"time"
)

const (
	PrivateAccessSecrete = `bCBsb3ZlIGJlZXJz` //此处可换成特定加密私钥，比如rsa生成的私钥信息
	TokenExpireTime      = 12
)

func GenJwtToken(user rbac.AuthUser) (string, error) {
	claims := model.PerchToken{
		UserUID:  user.UserUID,
		UserGID:  user.UserUID,
		UserName: user.UserName,
		Status:   user.UserStatus,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireTime * time.Hour).Unix(),
			//	Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(PrivateAccessSecrete))
}

func VerifyToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(PrivateAccessSecrete), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ParseJwtToken(tokenStr string) (model.PerchToken, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &model.PerchToken{}, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(PrivateAccessSecrete), nil

	})
	if err != nil {
		return model.PerchToken{}, err
	}
	if aesUser, ok := token.Claims.(*model.PerchToken); ok && token.Valid {
		return *aesUser, nil
	}
	return model.PerchToken{}, errors.New("Parse Token Failed!!!")
}
