package network

import (
	"fmt"
	"testing"
)

func TestGetDeviceIPAddress(t *testing.T) {
	ips, err := GetDeviceIPAddress()
	if err != nil {

	}
	fmt.Println(ips)
}
