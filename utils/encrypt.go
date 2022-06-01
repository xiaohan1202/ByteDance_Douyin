package utils

import (
	"crypto/sha256"
)

//随机加密盐对字符串进行SHA256加密，返回加密后哈希值
func Encrypt(s string, salt []byte) []byte {
	h := sha256.New()
	h.Write([]byte(s))
	h.Write(salt)
	return h.Sum(nil)
}
