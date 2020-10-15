package main

import (
	"net/http"
	"perch/api/spider_api"
	database "perch/database/mysql"
	"perch/web/service"
)

func main() {
	service.WebService{

		Name: "plat-spider",

		Router: []service.WebRouter{
			{RouterPath: "/spiders", RouterHandlerFunc: spider_api.CreateCollySpider, RouterMethod: http.MethodPost},
		},
		InitFunc: []func() error{
			database.InitMySQLDB,
		},
	}.WebServiceStart()

}
