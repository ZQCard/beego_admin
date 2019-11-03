package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"gopkg.in/gomail.v2"
	"log"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"sync"
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

func SendEmail(subject string, emails []string, content string) (success int, fail int, unreachable int) {
	// 多goroutine发送邮件
	var wg sync.WaitGroup
	host := beego.AppConfig.String("mail_host")
	port := MustInt(beego.AppConfig.String("mail_port"))
	username := beego.AppConfig.String("mail_username")
	password := beego.AppConfig.String("mail_password")
	for _,email := range emails{
		// 符合邮箱规则的发送邮件
		if VerifyEmailFormat(email) {
			wg.Add(1)
			go func(emailAddress string) {
				defer wg.Done()
				m := gomail.NewMessage()
				m.SetHeader("From", username)
				m.SetHeader("To", emailAddress)
				m.SetHeader("Subject", subject)
				m.SetBody("text/html", content)

				d := gomail.NewDialer(host, port, username, password)

				// 发送邮件
				if err := d.DialAndSend(m); err != nil {
					fail++
					// 发送邮件失败
					logs.Error("邮件发送失败", err, " email=", email)
				}else {
					success++
				}
			}(email)
		}else {
			unreachable++
		}
	}
	wg.Wait()
	return success, fail, unreachable
}

// 获取excel中的内容
func GetExcelContent(filePath string) [][]string {
	var content [][]string
	// 读取excel文件
	excelFile,err := excelize.OpenFile(filePath)
	if err != nil {
		log.Fatal("读取excel文件失败：", err)
	}
	// 迭代数据
	rows, err := excelFile.Rows("Sheet1")

	for rows.Next()  {
		row, err := rows.Columns()
		if err != nil{
			log.Fatal("读取excel数据失败：", err)
		}
		content = append(content, row)
		// 获取每个单元格内容
		// for _, colCell := range row {
		// 	 fmt.Println(colCell, "\t")
		// }
	}
	return content
}

func SetExcelContent(filename string, title *[]interface{}, content *[][]interface{}) (string, error) {
	f := excelize.NewFile()
	// 创建一个工作表
	//  err := f.SetSheetRow("Sheet1", "B6", &[]interface{}{"1", nil, 2})
	index := f.NewSheet("Sheet1")
	// 生成第一行数据
	f.SetSheetRow("Sheet1", "A1", title)
	// 依次传入数据 行+2
	for k, v := range *content{
		rowNum := k + 2
		err := f.SetSheetRow("Sheet1", "A" + strconv.Itoa(rowNum), &v)
		if err != nil {
			logs.Error("写入excel数据失败", err)
			return "", err
		}
	}
	f.SetActiveSheet(index)
	// 根据路径保存文件
	path := "static/example/" + filename + ".xlsx"
	err := f.SaveAs(path)
	if err != nil {
		log.Fatal("导出excel文件失败", err)
		return "", err
	}
	return path, nil
}

// excel日期字段格式化 yyyy-mm-dd
func ConvertToFormatDay(excelDaysString string)string{
	// 2006-01-02 距离 1900-01-01的天数
	baseDiffDay := 38719 //在网上工具计算的天数需要加2天，什么原因没弄清楚
	curDiffDay := excelDaysString
	b,_ := strconv.Atoi(curDiffDay)
	// 获取excel的日期距离2006-01-02的天数
	realDiffDay := b - baseDiffDay
	//fmt.Println("realDiffDay:",realDiffDay)
	// 距离2006-01-02 秒数
	realDiffSecond := realDiffDay * 24 * 3600
	//fmt.Println("realDiffSecond:",realDiffSecond)
	// 2006-01-02 15:04:05距离1970-01-01 08:00:00的秒数 网上工具可查出
	baseOriginSecond := 1136185445
	resultTime := time.Unix(int64(baseOriginSecond + realDiffSecond), 0).Format("2006-01-02")
	return resultTime
}

// sha1加密字符串
func Sha1String(data string) string {
	sha1 := sha1.New()
	sha1.Write([]byte(data))
	return hex.EncodeToString(sha1.Sum([]byte(nil)))
}
