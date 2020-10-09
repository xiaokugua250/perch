package auth

import (
	"fmt"
	"perch/web/model"
	"testing"
)

func TestGenJwtToken(t *testing.T) {
	var user model.User
	user.ID = 1
	user.UserStatus = 1
	user.UserName = "lixx"
	token, err := GenJwtToken(user)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("token is :", token)
}

func TestVerifyToken(t *testing.T) {
	token, err := VerifyToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGl4eCIsInVzZXJfcm9sZXMiOm51bGwsInN0YXR1cyI6MSwiZXhwIjoxNTk5MjI4Njc0LCJpc3MiOiJ0ZXN0In0.JubIMnaKSL96vTmPGUSl2zrNseYjJMXD7VF8buLY2jo")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
}
func TestParseJwtToken(t *testing.T) {
	token, err := ParseJwtToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcl9uYW1lIjoibGl4eCIsInVzZXJfcm9sZXMiOm51bGwsInN0YXR1cyI6MSwiZXhwIjoxNTk5MjI4Njc0LCJpc3MiOiJ0ZXN0In0.JubIMnaKSL96vTmPGUSl2zrNseYjJMXD7VF8buLY2jo")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%#v", token)
}
