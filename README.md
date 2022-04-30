# alioss

- golang 写的小工具
- 主要用来上传备份文件到阿里云oss以及删除阿里云oss备份文件


## config.ini
```
[oss]
AccessKeyId=""
AccessKeySecret=""
OSSEndpoint="https://oss-cn-beijing.aliyuncs.com"
OSSBucket="backup"
OSSDir="gitlab"
## 开始时间
StartTime="2020-01-01"
# 删除oss备份，2个月前的，请写"-2"
MonthsAgo=-3

[client]
# 删除日志，7天前的日志，请写"-7"
LogsDaysAgo=-7
BackupPath = "/data/backups/"
LogPath = "/data/back-logs/"
  
```

## 使用

```golang
package main

import (
    alioss "github.com/chaiyd/alioss/src"
)
  
func main() {
    alioss.UploadFile()
    alioss.OssDelFile()
}
```
