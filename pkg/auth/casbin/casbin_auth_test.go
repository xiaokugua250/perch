package auth

import (
	"fmt"


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



func TestCasbinAcm_CasbinAccessWithDB1(t *testing.T) {

	DBConfig := "root:mysqladmin@tcp(localhost:3306)/morty_db?charset=utf8mb4&parseTime=True&loc=Local"
	//dsn := DBConfig

	MysqlDb, err := gorm.Open(mysql.Open(DBConfig), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	casbinEnforcer, err := CasbinInit("./casbin.conf")
	if err != nil{
		log.Fatalln("error iswwww ",err)
	}


	request := CasbinSpecRequest{
		Subject: "liangdu",
		Domain:  "z-gour.com/*",
		Object:  "write",
	//	Actions: []string{"read"},
	}
	pass ,err :=		casbinEnforcer.CasbinAccessWithDB(MysqlDb, request)
	if err!= nil{
		log.Fatalln("error is----",err)
	}
	fmt.Println("====",pass)
}
