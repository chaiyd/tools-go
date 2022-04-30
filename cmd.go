package main

import (
	// "github.com/chaiyd/aliyun-tools/"
	"github.com/chaiyd/aliyun-tools/alilog"
)

func main() {
	alioss.UploadFile()
	alioss.OssDelFile()
	alilog.SendLog()
}
