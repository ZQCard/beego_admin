package utils

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/astaxie/beego"
)

var (
	endPoint string
	accessKeyId string
	accessKeySecret string
	bucketNameDefault string
)
// 定义进度条监听器。
type OssProgressListener struct {
}

// 定义进度变更事件处理函数。
func (listener *OssProgressListener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		fmt.Printf("Transfer Started, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferDataEvent:
		fmt.Printf("\rTransfer Data, ConsumedBytes: %d, TotalBytes %d, %d%%.",
			event.ConsumedBytes, event.TotalBytes, event.ConsumedBytes*100/event.TotalBytes)
	case oss.TransferCompletedEvent:
		fmt.Printf("\nTransfer Completed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	case oss.TransferFailedEvent:
		fmt.Printf("\nTransfer Failed, ConsumedBytes: %d, TotalBytes %d.\n",
			event.ConsumedBytes, event.TotalBytes)
	default:
	}
}

func init()  {
	// Endpoint以杭州为例，其它Region请按实际情况填写。
	endPoint = beego.AppConfig.String("end_point")
	// 阿里云主账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM账号进行API访问或日常运维，请登录 https://ram.console.aliyun.com 创建RAM账号。
	accessKeyId = beego.AppConfig.String("aliyun_oss_access_key")
	accessKeySecret = beego.AppConfig.String("aliyun_oss_access_secret")
	bucketNameDefault = beego.AppConfig.String("aliyun_oss_bucket")
}

func checkBucket(bucketName string) string {
	if bucketName == "" {
		bucketName = bucketNameDefault
	}
	return bucketName
}

// 创建OSS实例对象
func NewOssClient() (client *oss.Client, err error) {
	client, err = oss.New(endPoint, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// 创建Bucket
func CreateBucket(bucketName string) (err error) {
	client, err := NewOssClient()
	if err != nil{
		return err
	}
	bucketName = checkBucket(bucketName)
	fmt.Println(bucketName)
	// 创建存储空间。
	err = client.CreateBucket(bucketName)
	if err != nil {
		return err
	}
	return nil
}

// 删除bucket
func DeleteBucket(client *oss.Client, bucketName string) (err error) {
	bucketName = checkBucket(bucketName)
	// 删除存储空间。
	err = client.DeleteBucket(bucketName)
	if err != nil {
		return err
	}
	return nil
}
// 获取bucket
func GetBucket(bucketName string) (bucket *oss.Bucket,err error) {
	bucketName = checkBucket(bucketName)
	client, err := NewOssClient()
	if err != nil{
		return nil, err
	}
	bucket,err = client.Bucket(bucketName)
	if err != nil{
		return nil, err
	}
	return bucket, nil
}

// 上传文件
func UploadOssFile(bucketName, objectName, filePath string) (url string, err error) {
	bucket, err := GetBucket(bucketName)
	if err != nil{
		return "", err
	}

	err = bucket.PutObjectFromFile(objectName, filePath, oss.Progress(&OssProgressListener{}))
	if err != nil{
		return "", err
	}
	url = "//"+bucketNameDefault+"."+endPoint+"/"+objectName
	return url, nil
}

// 下载文件
func DownloadOssFile(bucketName, objectName, downloadedFileName string) (err error) {
	bucket, err := GetBucket(bucketName)
	if err != nil{
		return err
	}
	err = bucket.GetObjectToFile(objectName, downloadedFileName)
	if err != nil {
		return err
	}
	return nil
}

// 删除文件
func DeleteOssFile(bucketName, objectName string) (err error) {
	bucket, err := GetBucket(bucketName)
	if err != nil{
		return err
	}

	err = bucket.DeleteObject(objectName)
	if err != nil{
		return err
	}
	return nil
}