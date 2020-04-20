package pkg

import (
	"fmt"
	"testing"
)

func TestBase64Encoidng2Str(t *testing.T) {
	str := Base64Str2Encoding(KEY)
	fmt.Println("--->", str)
	str1 := Base64Encoidng2Str(str)
	fmt.Println(str1)
}
