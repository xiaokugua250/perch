// go test -v proxy_basic_test.go proxy_basic.go 
package proxy

import (
	"fmt"
	"log"
	"testing"
)


// go test -v -test.run  TestProxyServerSetup proxy_basic_test.go proxy_basic.go 
func TestProxyServerSetup(t *testing.T) {
	var (
		err error
	)
	fmt.Println("=============")
	err = ServerSetup()
	if err!= nil{
		log.Fatalln(err)
	}
	fmt.Printf("====%s","A")

}


func TestProxyServerUpWithOptions(t *testing.T){

	var (
		err error
	)
	err=ServerSetup()
	if err!= nil{
		log.Fatalln(err)
	}


}