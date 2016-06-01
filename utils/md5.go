package utils

import (
	"crypto/md5"
	"math/rand"
	"time"

	"encoding/hex"

	"github.com/astaxie/beego"
)

// GetRandomString return of Length lenNum Random string
func GetRandomString(lenNum int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// `i < len` will lead to the wrong --> `[r.Intn(len(str))]`
	for i := 0; i < lenNum; i++ {
		result = append(result, bytes[r.Intn(len(str))])
	}
	return string(result)
}

// GetSalt return the Salt
func GetSalt() string {
	len, err := beego.AppConfig.Int("lensalt")
	if err == nil {
		return GetRandomString(len)
	}
	return ""
}

// MD5 MD5 the Password + Salt
func MD5(text string) string {
	ctx := md5.New()
	ctx.Write([]byte(text))
	return hex.EncodeToString(ctx.Sum(nil))
}
