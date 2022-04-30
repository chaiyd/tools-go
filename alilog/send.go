package alilog

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/signal"
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
	ch := make(chan os.Signal)
	signal.Notify(ch, os.Kill, os.Interrupt)
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
	for {
		// line, err := reader.ReadString('\n')
		line, _, err := reader.ReadLine()
		// fmt.Printf("这一行的内容是：%s\n", line)
		if err == io.EOF {
			time.Sleep(time.Second * 1)
			continue
		}
		if err != nil && err != io.EOF {
			fmt.Println("读取错误", err)
			break
		}

		log := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"content": "test", "content2": fmt.Sprintf("%s\n", line)})
		producerInstance.SendLog(LOGProject, LOGLogstore, LOGTopic, LOGSource, log)
		producerInstance.SafeClose()

	}

}
