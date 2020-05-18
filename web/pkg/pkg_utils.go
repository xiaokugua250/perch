package pkg

import (
	"encoding/base64"
)

/**
将字符串进行base64 加密
*/
func Base64Str2Encoding(str string) string {

	return base64.StdEncoding.EncodeToString([]byte(str))

}

/**
base64 解密
*/
func Base64Encoidng2Str(encodedStr string) string {

	rawdata, _ := base64.StdEncoding.DecodeString(encodedStr)
	return string(rawdata)
}
