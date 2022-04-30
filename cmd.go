package main

import (
	"github.com/chaiyd/alioss/alioss"
	"github.com/chaiyd/aliyun-tools/alilog"
)

func main() {
	// c := make(chan os.Signal)
	// //监听所有信号
	// signal.Notify(c)
	// //阻塞直到有信号传入
	// fmt.Println("启动")
	// s := <-c
	// fmt.Println("退出信号", s)
	alioss.UploadFile()
	alioss.OssDelFile()
	alilog.AliSendLog()

}
