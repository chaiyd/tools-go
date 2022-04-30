package alilog

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/aliyun/aliyun-log-go-sdk/producer"
	"github.com/chaiyd/aliyun-tools/api"
)

func sendLog() {

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
	LOGFile := fmt.Sprint(cfg.Section("server").Key("LOGFileName"))

	f, err := os.Open(LOGFile)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}
		if line == "" {
			time.Sleep(time.Second * 1)
			continue
		}
		log := producer.GenerateLog(uint32(time.Now().Unix()), map[string]string{"content": "test", "content2": fmt.Sprintf("%v", line)})
		producerInstance.SendLog(LOGProject, LOGLogstore, LOGTopic, "", log)
		producerInstance.SafeClose()

	}

}
