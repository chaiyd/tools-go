package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	tools "github.com/chaiyd/tools-go/api"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
)

// env
//正式环境：prod，测试环境：test
var env = "test"
var ip = getip()
var CadvisorPort = "38098"
var CadvisorUrl = "http://" + ip + ":" + CadvisorPort + "/metrics"

// json type
//type(流量类型，1:下载,2:上传)

type Data struct {
	DockerName    string
	DockerId      string
	DockerNetwork string
	DockerIP      string
	Type          string
	Bytes         decimal.Decimal
	Timestamp     string
}

type Error struct {
	Error      string
	Msg        string
	DockerName string
	DockerId   string
}

func Write(data []byte) error {
	path := "/data/network/" + env + "/"
	time := time.Now().Format("2006-01-02")
	f, err := os.OpenFile(path+ip+"-"+time+".log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	_, err = f.Write(data)
	return err
}

func getip() string {
	ip, err := tools.GetIP()
	if err != nil {
		log.Println("获取ip出错", err)
	}
	return ip
}

// 流量
func Receive() {
	var network = "bridge"

	for _, container := range tools.ContainersList() {
		response, err := http.Get(CadvisorUrl)
		if err != nil {
			log.Println(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
		}

		networks := container.NetworkSettings
		networksType, err := json.Marshal(networks)
		if err != nil {
			log.Println(err)
		}
		networksStr := string(networksType)

		// 判断默认docker inspect network bridge 网卡是否存在
		value := gjson.Get(networksStr, "Networks."+network)
		if !value.Exists() {
			containerName := strings.Trim(container.Names[0], "/")
			dataerr := Error{
				Error:      "err",
				Msg:        "网卡不存在",
				DockerName: containerName,
				DockerId:   container.ID,
			}
			datajson, _ := json.Marshal(dataerr)

			datajsonStr := string(datajson)
			err = Write([]byte(datajsonStr + "\n"))
			if err != nil {
				log.Println(err)
			}
			// fmt.Println("关闭资源,网卡不存在11111", response.Body.Close())
			continue
		} else {
			r_source := regexp.MustCompile("container_network_receive_bytes_total"+
				"(.*)"+container.ID+"(.*)"+
				"eth0"+"(.*)").FindAllStringSubmatch(string(body), -1)
			// fmt.Println(r_source)

			r_source_string := strings.Split(r_source[0][0], " ")
			r_source_time := r_source_string[len(r_source_string)-1]
			r_source_num_str := r_source_string[len(r_source_string)-2]
			r_source_num, _ := decimal.NewFromString(r_source_num_str)
			containerName := strings.Trim(container.Names[0], "/")
			containerIP := container.NetworkSettings.Networks[network].IPAddress

			data := Data{
				DockerName:    containerName,
				DockerId:      container.ID,
				DockerNetwork: network,
				DockerIP:      containerIP,
				Type:          "1",
				Bytes:         r_source_num,
				Timestamp:     r_source_time,
			}
			datajson, err := json.Marshal(data)
			if err != nil {
				log.Println(err)
			}

			datajsonStr := string(datajson)
			err = Write([]byte(datajsonStr + "\n"))
			if err != nil {
				log.Println(err)
			}
			// fmt.Println("关闭资源333333", response.Body.Close())
			continue
		}
	}
}

func Transmit() {
	var network = "bridge"
	for _, container := range tools.ContainersList() {
		response, err := http.Get(CadvisorUrl)
		if err != nil {
			log.Println(err)
		}
		defer response.Body.Close()
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Println(err)
		}

		networks := container.NetworkSettings
		networksType, err := json.Marshal(networks)
		if err != nil {
			log.Println(err)
		}
		networksStr := string(networksType)

		value := gjson.Get(networksStr, "Networks."+network)
		if !value.Exists() {
			containerName := strings.Trim(container.Names[0], "/")
			dataerr := Error{
				Error:      "err",
				Msg:        "网卡不存在",
				DockerName: containerName,
				DockerId:   container.ID,
			}
			datajson, _ := json.Marshal(dataerr)
			datajsonStr := string(datajson)
			err = Write([]byte(datajsonStr + "\n"))
			if err != nil {
				log.Println(err)
			}
			// fmt.Println("关闭资源,网卡不存在22222221", response.Body.Close())
			continue
		} else {
			t_source := regexp.MustCompile("container_network_transmit_bytes_total"+"(.*)"+container.ID+"(.*)"+
				"eth0"+"(.*)").FindAllStringSubmatch(string(body), -1)

			t_source_string := strings.Split(t_source[0][0], " ")
			t_source_time := t_source_string[len(t_source_string)-1]
			t_source_num_str := t_source_string[len(t_source_string)-2]
			t_source_num, _ := decimal.NewFromString(t_source_num_str)
			containerName := strings.Trim(container.Names[0], "/")
			containerIP := container.NetworkSettings.Networks[network].IPAddress

			data := Data{
				DockerName:    containerName,
				DockerId:      container.ID,
				DockerNetwork: network,
				DockerIP:      containerIP,
				Type:          "2",
				Bytes:         t_source_num,
				Timestamp:     t_source_time,
			}
			datajson, err := json.Marshal(data)
			if err != nil {
				log.Println(err)
			}
			datajsonStr := string(datajson)
			err = Write([]byte(datajsonStr + "\n"))
			if err != nil {
				log.Println(err)
			}
			// fmt.Println("关闭资源44444", response.Body.Close())
			continue
		}
	}
}

func main() {
	for {
		go func() {
			Receive()
			Transmit()
			defer tools.SystemExit()
		}()
		time.Sleep(time.Second * 1)
	}
}
