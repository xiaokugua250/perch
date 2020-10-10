package database

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var GormDB *sql.DB

func InitDataBase() error {
	var (
		err error
	)

	MySQLDB, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置

	}), &gorm.Config{
		// GORM perform single create, update, delete operations in transactions by default to ensure database data integrity
		// You can disable it by setting `SkipDefaultTransaction` to true
		//SkipDefaultTransaction:false,
		// NamingStrategy tables, columns naming strategy
		//	NamingStrategy schema.Namer
		// Logger
		//	Logger:logger.New(),
		// NowFunc the function to be used when creating a new timestamp
		//NowFunc func() time.Time
		// DryRun generate sql without execute
		//DryRun bool
		// PrepareStmt executes the given query in cached statement
		//PrepareStmt bool
		// DisableAutomaticPing
		//DisableAutomaticPing bool
		// DisableForeignKeyConstraintWhenMigrating
		//	DisableForeignKeyConstraintWhenMigrating bool
		// AllowGlobalUpdate allow global update
		//AllowGlobalUpdate bool

		// ClauseBuilders clause builder
		//ClauseBuilders map[string]clause.ClauseBuilder
		// ConnPool db conn pool
		//	ConnPool ConnPool
		// Dialector database dialector
		//	Dialector
		// Plugins registered plugins
		//	Plugins map[string]Plugin

		//	callbacks  *callbacks
		//	cacheStore *sync.Map
	})
	GormDB, err = MySQLDB.DB()
	if err != nil {
		return err
	}
	GormDB.SetMaxIdleConns(10)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	GormDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	GormDB.SetConnMaxLifetime(time.Hour)

	return nil

}
