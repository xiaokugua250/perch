package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MySQL_DB *gorm.DB
)

//初始化数据库
func InitMySQLDB() error {
	DBConfig := "genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := DBConfig
	var (
		err error
	)
	MySQL_DB, err = gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
	return err
}
