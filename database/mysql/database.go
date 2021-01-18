package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	MysqlDb *gorm.DB
)

//初始化数据库
func InitMySQLDB() error {
	DBConfig := "genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := DBConfig
	var (
		err error
	)
	MysqlDb, err = gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
	if err != nil {
		return err
	}
	MysqlDb.Logger = logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)

	//MySQL_DB = MySQL_DB.Debug()

	return err
}

//初始化数据库
func InitMySQLDBWithConig(config interface{}) error {

	fmt.Print(config)
	DBConfig := "genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := DBConfig
	var (
		err error
	)
	MysqlDb, err = gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
	return err
}
