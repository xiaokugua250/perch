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
	err = ServerInit()
	if err != nil {
		log.Fatalln(err)
	}
	err = ServerSetup()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("====%s", "A")

}

func TestProxyServerUpWithOptions(t *testing.T) {

	var (
		err error
	)

	err = ServerSetup()
	if err != nil {
		log.Fatalln(err)
	}

}

func TestServerSetup(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ServerSetup(); (err != nil) != tt.wantErr {
				t.Errorf("ServerSetup() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
