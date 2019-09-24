package main

import "fmt"

func main ()  {
	var a int
	a = 10
	a = 20    //赋值可以使用n次
	fmt.Println("a = ", a)

	// :=, 自动推导类型,先声明变量b,在给b赋值为20
	//自动推导,只能使用一次,用于初始化那次
	b := 20
	fmt.Println("b = ", b)
	//b := 30    前面已有变量b.不是再新建一个变量b
	b = 30
	fmt.Println("b2 = ", b)
}

//返回结果
/*
a =  20
b =  20
b2 =  30

 */