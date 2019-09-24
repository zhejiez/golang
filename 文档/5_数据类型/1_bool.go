package main

import "fmt"

//bool型
func main()  {
	//1, 声明变量,没有初始化,零值(初始值)为false
	var a bool
	fmt.Println("a = ", a)
	a = true
	fmt.Println("a = ", a)

	//2m 自动推导类型
	var b = false
	fmt.Println("b = ", b)

	c := false
	fmt.Println("c =", c)
}

//返回结果
/*
a =  false
a =  true
b =  false
c = false
 */