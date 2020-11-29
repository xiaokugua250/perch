package service

import (
	"fmt"
	"testing"
)

type FuncWithConfigs struct {
	InitFunc []func(config interface{}) error
}

func awith(a interface{}) error {
	fmt.Println(a)
	return nil
}
func bwith(a interface{}) error {
	fmt.Println(a)
	return nil
}
func cwith(a interface{}) error {
	fmt.Println(a)
	return nil
}
func TestWebServer_StartServer(t *testing.T) {
	fmt.Print("xasxsxs")
	var initFunc []func(interface{}) error
	initFunc = append(initFunc, awith)
	initFunc = append(initFunc, bwith)
	initFunc = append(initFunc, cwith)
	funs := &FuncWithConfigs{InitFunc: initFunc}

	for key, fu := range funs.InitFunc {
		fmt.Printf("key is %+v,%+v\n", key, fu("a"))
	}
}
