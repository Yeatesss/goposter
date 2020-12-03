package file

import (
	"fmt"
	"os"
	"sync"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

var (
	o    *OssFile
	lock sync.Mutex
)

type OssFile struct {
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
}

func NewOssFile() (oss *OssFile) {
	if nil == o {
		lock.Lock()
		if nil == o {
			o = &OssFile{
				Endpoint:        viper.GetString("oss.endpoint"),
				AccessKeyId:     viper.GetString("oss.access_key_id"),
				AccessKeySecret: viper.GetString("oss.access_key_secret"),
				BucketName:      viper.GetString("oss.bucket_name"),
			}
			lock.Unlock()
		} else {
			lock.Unlock()
		}
	}
	return o
}

func (o *OssFile) PrintTest() {
	fmt.Println("oss")
}

func (o *OssFile) Save(savePath, fileName string) (err error) {
	client, err := oss.New(o.Endpoint, o.AccessKeyId, o.AccessKeySecret)
	if err != nil {
		return
	}
	// 获取存储空间。
	bucket, err := client.Bucket(o.BucketName)
	if err != nil {
		return
	}
	src, _ := os.Open(viper.GetString("img_tmp_dir") + fileName)
	err = bucket.PutObject(savePath+fileName, src)
	if err != nil {
		return err
	}
	return
}
