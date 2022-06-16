package api

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

// 一些常用方法

// 关闭程序
func SystemExit() {
	c := make(chan os.Signal, 1)
	//监听信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		s := <-c
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			log.Println("退出程序：", s)
			os.Exit(0)
		default:
			return
		}
	}

}

// 获取本机内网ip
func GetIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				return
				// fmt.Println("test", ipnet.IP.String())
			}
		}
	}
	return
}

// 获取公网ip
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

// 获取本机hostname
func GetHOSTNAME() (hostname string, err error) {
	hostname, err = os.Hostname()
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

// func main() {
// 	ip, err := GetOutBoundIP()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(ip)

// 	hostname, err := GetHOSTNAME()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println("hostname", hostname)

// }
