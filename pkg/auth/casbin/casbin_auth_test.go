package auth

import (
	"fmt"

	"github.com/casbin/casbin/v2"
	_ "github.com/go-sql-driver/mysql"

	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func TestCasbinInit(t *testing.T){

	casbinEnforcer, err := CasbinInit("./casbin.conf")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("----",casbinEnforcer)

}


func TestCasbinAcm_CasbinAccessWithDB(t *testing.T) {

	DBConfig := "root:mysqladmin@tcp(localhost:3306)/morty_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := DBConfig

	MysqlDb, err := gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	casbinEnforcer, err := CasbinInit("./casbin.conf")
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
			}{db: MysqlDb, request: struct {
				Subject string
				Domain  string
				Object  string
				Actions []string
			}{Subject: "a", Domain: "z-gour.com", Object: "resources", Actions: []string{"read"}}},
			true,
			false,
		},
	}
	fmt.Println("-----")
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



func TestCasbinAcm_CasbinAccessWithDB1(t *testing.T) {

	DBConfig := "root:mysqladmin@tcp(localhost:3306)/morty_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := DBConfig

	MysqlDb, err := gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	casbinEnforcer, err := CasbinInit("./casbin.conf")

	request := CasbinSpec{
		Subject: "liangdu",git 
		Domain:  "z-gour.com",
		Object:  "resource",
		Actions: []string{"read"},
	}
	pass ,err :=		casbinEnforcer.CasbinAccessWithDB(MysqlDb, request)
	if err!= nil{
		log.Fatalln(err)
	}
	fmt.Println("====,",pass)
}
