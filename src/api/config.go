package api

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func LoadConfig() *ini.File {
	// 读取配置文件
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	path = path[:index]
	// fmt.Print("path:", path)
	cfg, err := ini.Load(path + "/conf/config.ini")
	if err != nil {
		fmt.Println("文件读取错误:", err)
		os.Exit(1)
	}
	return cfg
}
