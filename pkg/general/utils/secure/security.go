package secure

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
)

const Charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789~=+%^*/()[]{}/!@#$?|" //原始随机字符串

/**
生成随机字符串
*/
func GenerateRandomeStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = Charset[rand.Intn(len(Charset))]
	}
	return string(b)
}

/**
生成字符串MD5值
*/
func GenerateMd5Hash(str string) string {
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

/**
生成字符串SSH1 加密
*/
func GenerateSHA1Hash(str string) string {
	hasher := sha1.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}
