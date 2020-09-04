package model

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type PEXToken struct {
	UserID             int      `json:"id"`
	UserName           string   `json:"user_name"` // 用户名
	UserRoles          []string `json:"user_roles"`
	Status             int      `json:"status"` // 用户状态
	jwt.StandardClaims          // 标准JWT包含的Token内容
}

type User struct {
	ID              int       `json:"id" gorm:"column:id;type:int(11);not null;primary_key"`
	UserUID         string    `json:"useruid" gorm:"column:useruid;type:int(10);not null"`
	UserGID         string    `json:"usergid" gorm:"column:usergid;type:int(10);not null"`
	UserName        string    `json:"username" gorm:"column:username;type:varchar(128);not null"`
	UserPasswd      string    `json:"password" gorm:"column:password;type:varchar(128);not null"`
	UserEmail       string    `json:"email" gorm:"column:email;type:varchar(128);not null"`
	UserAvatar      string    `json:"user_avatar" gorm:"column:status;type:varchar(128);not null"`
	UserLastLogin   time.Time `json:"user_last_login" gorm:"column:lastlogin;type:varchar(128);not null"`
	UserStatus      int       `json:"user_status" gorm:"column:status;type:tinyint(4);not null"`
	UserCreatedAt   time.Time `json:"created_at" gorm:"column:created_at;type:bigint(20);not null"`
	UserUpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at;type:bigint(20);not null"`
	UserDescription string    `json:"user_description" gorm:"column:description;type:varchar(128);not null"`
}
