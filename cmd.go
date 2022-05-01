package main

import (
	"github.com/chaiyd/aliyun-tools/alilog"
	"github.com/chaiyd/aliyun-tools/alioss"
)

func main() {
	alioss.UploadFile()
	alioss.OssDelFile()
	alilog.AliSendLog()

}
