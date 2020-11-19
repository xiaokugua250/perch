package cookies

import (
	"log"
	"testing"
)

func TestScannLocalCookies(t *testing.T) {
	err :=ScannLocalCookies("","","")
	if err!= nil{
		log.Println(err)
	}

}
