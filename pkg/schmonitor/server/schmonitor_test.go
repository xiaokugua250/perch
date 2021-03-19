package server

import (
	"log"
	"testing"
)

func TestServerSetupWithOpt(t *testing.T) {

	var options SetupOptions
	options.IP = "127.0.0.1"
	options.Port = "9999"
	options.Protocols = []int{1, 2, 3}
	if err := ServerSetupWithOpt(options); err != nil {
		log.Println(err)
	}
}
