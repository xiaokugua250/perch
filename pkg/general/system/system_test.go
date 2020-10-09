package system

import (
	"fmt"
	"testing"
)

func TestSignaleHander(t *testing.T) {
	SignaleHander()
}

func TestExecuteCmdWithParams(t *testing.T) {
	err, result := ExecuteCmdWithParams("java", "-help")
	fmt.Println("err is ,result is", err, result)

}
