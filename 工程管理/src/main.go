package main

import (
	"calc"
	"fmt"
	"structural"
)

func init() {
	fmt.Println("main init函数")
}

//运行命令:go run main.go test.go
func main() {
	//包名.函数名
	structural.MyFunc()
	a := calc.Add(10, 20)
	fmt.Println("不同目录调用:a = ", a)

	fmt.Println("不同目录调用:r =", calc.Minus(10, 5))
	//包名.结构体里类型名
	var s structural.Stu
	s.Id = 666
	fmt.Println("s = ", s)
}
