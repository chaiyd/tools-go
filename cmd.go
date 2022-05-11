package main

import (
	"github.com/chaiyd/tools-go/aliyun"
)

func main() {
	aliyun.UploadFile()
	aliyun.OssDelFile()
	aliyun.AliSendLog()
}
