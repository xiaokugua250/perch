package auth

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

func TestCasbinAcm_CasbinAccessWithDB(t *testing.T) {

	DBConfig := "genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := DBConfig
	var (
		err error
	)
	MysqlDb, err := gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	casbinEnforcer, err := CasbinInit("../../")
	if err != nil {
		log.Fatalln(err)
	}
	type fields struct {
		SyncedEnforcer *casbin.SyncedEnforcer
	}
	type args struct {
		db      *gorm.DB
		request CasbinSpec
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"a",
			struct{ SyncedEnforcer *casbin.SyncedEnforcer }{SyncedEnforcer: casbinEnforcer.SyncedEnforcer},
			struct {
				db      *gorm.DB
				request CasbinSpec
			}{db: MysqlDb, request: nil},
			true,
			false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			casbinAcm := &CasbinAcm{
				SyncedEnforcer: tt.fields.SyncedEnforcer,
			}
			got, err := casbinAcm.CasbinAccessWithDB(tt.args.db, tt.args.request)
			if (err != nil) != tt.wantErr {
				t.Errorf("CasbinAccessWithDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CasbinAccessWithDB() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCasbinAcm_CasbinAccessWithDB2(t *testing.T) {
	CasbinAcm, err := CasbinInit("../")
	if err != nil {
		log.Fatalln(err)
	}
	DBConfig := "genuser:mysql123Admin@@tcp(172.16.171.84:3306)/morty?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := DBConfig

	MysqlDb, err := gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	requst := CasbinSpec{
		Subject: "",
		Domain:  "",
		Object:  "",
		Actions: nil,
	}
	pass, err := CasbinAcm.CasbinAccessWithDB(MysqlDb, requst)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(pass)
}
