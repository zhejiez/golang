package main

import (
	"calc"
	"fmt"
)
func init()  {
	fmt.Println("main init函数")
}

//运行命令:go run main.go test.go
func main()  {
	test()

	a := calc.Add(10, 20)
	fmt.Println("不同目录调用:a = ", a)

	fmt.Println("不同目录调用:r =", calc.Minus(10, 5))
}