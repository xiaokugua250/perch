package cookies

import (
	"fmt"
	"log"
	"testing"
)

func TestScannLocalCookies(t *testing.T) {
	err :=ScannLocalCookies("","","")
	if err!= nil{
		log.Println(err)
	}

}

func TestLoadCookies(t *testing.T) {
	cookies ,err := LoadCookies("windowscookies.json")
	if err!= nil{
		log.Println(err)
	}
	fmt.Println(cookies)
}
