# alioss

- golang 写的小工具

## alioss
- aliossdel
  - 用来删除alioss指定时间内的文件
- aliossupload
  - 上传指定目录文件到alioss
## alilog
  - 上传指定日志文件到alisls
  
---
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
LOGFile="/Users/videopls/Projects/aliyun-tools/alilog/log"
LOGSource="127.0.0.1"

# 删除本地日志，7天前的日志，请写"-7"
LogsDaysAgo=-7
BackupPath = "/data/backups/"
LogPath = "/data/back-logs/"
```

## 使用

```golang
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

```
