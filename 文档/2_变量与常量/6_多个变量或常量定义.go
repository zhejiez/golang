package main

import "fmt"

func main()  {
	//不同类型变量声明(定义)
	var (
		a int
		b float64
	)
	a, b = 10, 3.14
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//const i int = 10
	//const j float64  = 3.14
	const (
		i = 10//可以自动推导数据类型
		j = 3.14
	)
	fmt.Println("i = ", i)
	fmt.Println("j = ", j)
}

//返回结果
/*
a =  10
b =  3.14
i =  10
j =  3.14
 */