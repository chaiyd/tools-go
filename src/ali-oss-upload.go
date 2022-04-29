package alioss

import (
	"github.com/chaiyd/alioss/src/api"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func UploadFile() {

	cfg := api.LoadConfig()

	OSSEndpoint := fmt.Sprint(cfg.Section("oss").Key("OSSEndpoint"))
	AccessKeyId := fmt.Sprint(cfg.Section("oss").Key("AccessKeyId"))
	AccessKeySecret := fmt.Sprint(cfg.Section("oss").Key("AccessKeySecret"))
	// 创建OSSClient实例。
	client, err := oss.New(OSSEndpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	OSSBucket := fmt.Sprint(cfg.Section("oss").Key("OSSBucket"))
	// 填写Bucket名称，例如examplebucket。
	bucket, err := client.Bucket(OSSBucket)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	BackupPath := fmt.Sprint(cfg.Section("client").Key("BackupPath"))
	OSSDir := fmt.Sprint(cfg.Section("oss").Key("OSSDir"))

	files, err := ioutil.ReadDir(BackupPath)
	if err != nil {
		panic(err)
	}
	// 获取文件，并输出它们的名字
	for _, file := range files {
		println(file.Name())
		// 设置分片
		err = bucket.UploadFile(OSSDir+"/"+file.Name(),
			BackupPath+file.Name(),
			10*1024*1024, oss.Routines(10),
			oss.Checkpoint(true, file.Name()+".cp"))

		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		} else {
			fmt.Println("上传文件到oss成功:", file.Name())
		}

		// 删除备份文件
		err = os.Remove(BackupPath + file.Name())
		if err != nil {
			fmt.Println("删除文件失败:", err)
			os.Exit(-1)
		} else {
			fmt.Println("删除文件成功:" + file.Name())
		}
	}

	// 删除日志文件
	LogsDaysAgoStr := fmt.Sprint(cfg.Section("client").Key("LogsDaysAgo"))
	LogsDaysAgoInt, err := strconv.Atoi(LogsDaysAgoStr)
	if err != nil {
		fmt.Println("时间类型转换出错:", err)
		os.Exit(-1)
	}

	LogPath := fmt.Sprint(cfg.Section("client").Key("LogPath"))
	LogsDaysAgo := time.Now().AddDate(0, 0, LogsDaysAgoInt).Format("20060102")
	// fmt.Println(logsDaysago)
	logfile := LogPath + "gitlab-" + LogsDaysAgo + ".log"
	err = os.Remove(logfile)
	if err != nil {
		fmt.Println("删除日志文件失败:", err)
		os.Exit(-1)
	} else {
		fmt.Println("删除日志文件成功:" + logfile)
	}
}
