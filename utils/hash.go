package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

//	此包用于将字符串生成哈希的函数

func GenerateHash(s string) string {
	b := []byte(s)
	hash := sha256.New()
	hash.Write(b)
	hashed := hash.Sum(nil)
	//nil可以换成字节切片，实现追加功能
	hashedHex := hex.EncodeToString(hashed)
	//	此处将[]byte转换为16进制字符串了
	return hashedHex
}

// 相等就返回true
func CheckHash(s string, hash string) bool {
	hashed := GenerateHash(s)
	if fmt.Sprintf("%s", hashed) == fmt.Sprintf("%s", hash) {
		return true
	}
	return false
}
