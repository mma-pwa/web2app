package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

//@function: MD5V
//@description: md5加密
//@param: str []byte
//@return: string

func MD5V(str []byte, b ...byte) string {
	hash := md5.Sum(str)

	hash16 := hash[:16]
	md5Str := hex.EncodeToString(hash16)
	return md5Str[:16]
}

// Md5Crypt
// @给字符串生成md5
// @params str 需要加密的字符串
// @params salt interface{} 加密的盐
// @return str 返回md5码
func Md5Crypt(str string, salt ...interface{}) (CryptStr string) {
	if l := len(salt); l > 0 {
		slice := make([]string, l+1)
		str = fmt.Sprintf(str+strings.Join(slice, "%v"), salt...)
	}
	hash := md5.Sum([]byte(str))
	return fmt.Sprintf("%x", hash[:16])
}

func GetRandomString(n int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	var result []byte
	for i := 0; i < n; i++ {
		result = append(result, bytes[rand.Intn(len(bytes))])
	}
	return string(result)
}

func GetPassWord(userName, passWord, salt string) string {
	md5 := MD5V([]byte(userName + passWord))
	md5PassWord := Md5Crypt(md5, salt)
	return md5PassWord
}
