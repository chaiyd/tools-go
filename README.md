# alioss

- golang 写的小工具
- 主要用来上传备份文件到阿里云oss以及删除阿里云oss备份文件


## config.ini
```ini
[server]
AccessKeyId=""
AccessKeySecret=""
LOGEndpoint="cn-beijing.log.aliyuncs.com"
LOGProject="tx-ai-cloud"
LOGLogstore="txlog"
LOGTopic="test"

OSSEndpoint="https://oss-cn-beijing.aliyuncs.com"
OSSBucket="backup"
OSSDir="gitlab"
## 开始时间
StartTime="2020-01-01"
# 删除oss备份，2个月前的，请写"-2"
MonthsAgo=-3

[client]
# 上传阿里云sls
LOGFile="/Users/test.log"
LOGSource="127.0.0.1"

# 删除日志，7天前的日志，请写"-7"
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
