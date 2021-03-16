package proxy

import (
	"fmt"
	"log"
	"testing"
)



func TestProxyServerSetup(t *testing.T) {
	var (
		err error
	)
	err = ServerSetup()
	if err!= nil{
		log.Fatalln(err)
	}
	fmt.Printf("====%s","A")

}
