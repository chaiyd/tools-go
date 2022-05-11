# tools-go

- golang 写的小工具
- <https://github.com/chaiyd/tools-go.git>

## aliyun

- aliossdel
  - 用来删除alioss指定时间内的文件
- aliossupload
  - 上传指定目录文件到alioss
- alisls
  - 上传指定日志文件到alisls
  
## config.ini

```ini
[server]
AccessKeyId=""
AccessKeySecret=""
## alisls
LOGEndpoint="cn-beijing.log.aliyuncs.com"
LOGProject="log"
LOGLogstore="log"
LOGTopic="log"

## alioss
OSSEndpoint="https://oss-cn-beijing.aliyuncs.com"
OSSBucket="backup"
OSSDir="gitlab"
## 开始时间
OSSStartTime="2020-01-01"
# 删除oss备份，2个月前的，请写"-2"
OSSMonthsAgo=-3

[client]
# 上传阿里云sls
LOGFile="/Projects/tools-go/log"

# 删除本地日志，7天前的日志，请写"-7"
LogsDaysAgo=-7
BackupPath = "/data/backups/"
LogPath = "/data/back-logs/"
```

## 使用

```golang
package main

import (
  "github.com/chaiyd/tools-go/aliyun"
)

func main() {
  aliyun.UploadFile()
  aliyun.OssDelFile()
  aliyun.AliSendLog()
}

```
