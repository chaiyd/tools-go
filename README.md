# alioss

- golang 写的小工具
- 主要用来上传备份文件到阿里云oss以及删除阿里云oss备份文件

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
