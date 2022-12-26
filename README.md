# How Run 

# 一、这是啥
用来判断当前程序是怎么跑起来的，比如是源代码运行的，还是编译后运行的，是运行的单元测试，还是main方法运行的，是否泡在IDE比如GoLand中

# 二、安装 
```bash
go get -u github.com/golang-infrastructure/go-how-run
```

# 三、Example

```go
package main

import (
	"fmt"
	how_run "github.com/golang-infrastructure/go-how-run"
)

func main() {

	// 识别当前是运行的发布的二进制包还是从源代码运行
	runType, err := how_run.GetRunType()
	if err != nil {
		fmt.Println("GetRunType error: " + err.Error())
		return
	}
	fmt.Println(runType) // Output: SourceCode

	// 如果是从源代码运行，入口是啥，是main方法还是单元测试啥的
	sourceCodeRunType, err := how_run.GetSourceCodeRunType()
	if err != nil {
		fmt.Println("GetSourceCodeRunType error: " + err.Error())
		return
	}
	fmt.Println(sourceCodeRunType) // Output: main.go

	// 识别当前是否运行在IDE中
	ide, err := how_run.GetRunIDE()
	if err != nil {
		fmt.Println("GetRunIDE error: " + err.Error())
		return
	}
	fmt.Println(ide) // Output: GoLand

}
```
