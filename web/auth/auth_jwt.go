package auth

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"perch/web/model"

	"time"
)

const (
	PRIVATE_ACCESS_SECRETE = `bCBsb3ZlIGJlZXJz` //此处可换成特定加密私钥，比如rsa生成的私钥信息
	TOKEN_EXPIRE_TIME      = 12
)

func GenJwtToken(user model.AuthUser) (string, error) {
	claims := model.PEXToken{
		UserID:   user.ID,
		UserName: user.UserName,
		Status:   user.UserStatus,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TOKEN_EXPIRE_TIME * time.Hour).Unix(),
			//	Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(PRIVATE_ACCESS_SECRETE))
}

func VerifyToken(tokenString string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		//Make sure that the token method conform to "SigningMethodHMAC"
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(PRIVATE_ACCESS_SECRETE), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func ParseJwtToken(tokenStr string) (model.PEXToken, error) {

	token, err := jwt.ParseWithClaims(tokenStr, &model.PEXToken{}, func(token *jwt.Token) (i interface{}, e error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(PRIVATE_ACCESS_SECRETE), nil

	})
	if err != nil {
		return model.PEXToken{}, err
	}
	if aesUser, ok := token.Claims.(*model.PEXToken); ok && token.Valid {
		return *aesUser, nil
	}
	return model.PEXToken{}, errors.New("Parse Token Failed!!!")
}
