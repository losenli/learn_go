package main

import "fmt"

/**
# 编译命令
1、go build 直接编译当前项目
2、go build XXX.go 在GOPATH目录下寻找XXX.go并编译
3、go build -o XXX 指定编译完成后可执行程序的名字
4、go install 编译并保存到GOPATH/bin目录下
5、go run XXX.go 编译并执行XXX.go程序
 */

func main() {
	fmt.Println("Hello,World!")
}
