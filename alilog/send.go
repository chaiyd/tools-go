package alilog

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/chaiyd/aliyun-tools/api"
)

func AliSendLog() {

	cfg := api.LoadConfig()

	producerConfig := producer.GetDefaultProducerConfig()
	producerConfig.Endpoint = fmt.Sprint(cfg.Section("server").Key("LOGEndpoint"))
	producerConfig.AccessKeyID = fmt.Sprint(cfg.Section("server").Key("AccessKeyId"))
	producerConfig.AccessKeySecret = fmt.Sprint(cfg.Section("server").Key("AccessKeySecret"))
	producerInstance := producer.InitProducer(producerConfig)
	// ch := make(chan os.Signal)
	// signal.Notify(ch, os.Kill, os.Interrupt)
	producerInstance.Start() // 启动producer实例

	// 填写日志组名称。

	LOGProject := fmt.Sprint(cfg.Section("server").Key("LOGProject"))
	LOGLogstore := fmt.Sprint(cfg.Section("server").Key("LOGLogstore"))
	LOGTopic := fmt.Sprint(cfg.Section("server").Key("LOGTopic"))
	LOGFile := fmt.Sprint(cfg.Section("client").Key("LOGFile"))
	LOGSource := fmt.Sprint(cfg.Section("client").Key("LOGSource"))

	// fmt.Println("LOGFile:", LOGFile)
	f, err := os.Open(LOGFile)
	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	reader := bufio.NewReader(f)
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)

	for {

		s := <-c
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			producerInstance.SafeClose()
			fmt.Println("退出程序：", s)
			os.Exit(0)
		default:
			return
		}

		line, _, err := reader.ReadLine()
		// fmt.Printf("这一行是：", fmt.Sprintf("%s%v\n", line)
		if err == io.EOF {
			// fmt.Printf("数据读取完毕")
			time.Sleep(time.Second * 1)
			continue
		}
		if err != nil && err != io.EOF {
			fmt.Println("读取错误", err)
			os.Exit(-1)
		}
		log := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"content": fmt.Sprintf("%s\n", line)})
		producerInstance.SendLog(LOGProject, LOGLogstore, LOGTopic, LOGSource, log)

	}
	// fmt.Println("发送成功")
}
