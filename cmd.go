package main

import (
	alioss "github.com/chaiyd/alioss-tools/alioss"
	// alilog "github.com/chaiyd/alioss-tools/alilog"
)

func main() {
	alioss.UploadFile()
	alioss.OssDelFile()
	// alilog.SendLog()
}
