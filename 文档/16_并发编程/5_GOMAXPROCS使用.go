package main

import (
	"fmt"
	"runtime"
)

//设置运行使用CPU核数
func main() {
	n := runtime.GOMAXPROCS(8) //指定以8核运算
	fmt.Println(n)

	for {
		go fmt.Print(1)

		fmt.Print(0)
	}
}
