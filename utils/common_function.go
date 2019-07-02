package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gopkg.in/gomail.v2"
	"math/rand"
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

// 获取随机字符串
func RandString(len int) string {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := r.Intn(26)+65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func SendEmail(subject string, emails []string, content string) (success int, fail int) {
	host := beego.AppConfig.String("mail_host")
	port := MustInt(beego.AppConfig.String("mail_port"))
	username := beego.AppConfig.String("mail_username")
	password := beego.AppConfig.String("mail_password")
	for _,email := range emails{
		m := gomail.NewMessage()
		m.SetHeader("From", username)
		m.SetHeader("To", email)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", content)

		d := gomail.NewDialer(host, port, username, password)

		// Send the email to Bob, Cora and Dan.
		if err := d.DialAndSend(m); err != nil {
			fail++
			// 发送邮件失败
			logs.Error("邮件发送失败", err, " email=", email)
			fmt.Println(err)
		}else {
			success++
		}
	}
	return success, fail
}