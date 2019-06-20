package utils

import (
	"crypto/md5"
	"encoding/hex"
	"reflect"
	"regexp"
	"strconv"
	"time"
)

// 获取变量类型
func TypeOf(v interface{}) string {
	return reflect.TypeOf(v).String()
}

// 生成MD5字符串
func GenerateMD5String(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// 将字符串转换为int
func MustInt(s string) int {
	v,_ := strconv.Atoi(s)
	return v
}

// int64转换为int
func Int64ToInt(i int64) int {
	strInt64 := strconv.FormatInt(i, 10)
	num,_ := strconv.Atoi(strInt64)
	return num
}

// 将时间戳转换为年月日
func TimestampToDate(timestamp int) string {
	timestampString := strconv.Itoa(timestamp)
	timestamp64,_ := strconv.ParseInt(timestampString, 10, 64)
	dateTime := time.Unix(timestamp64, 0).Format("2006-01-02")
	return dateTime
}

// 验证邮箱
func VerifyEmailFormat(email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}
