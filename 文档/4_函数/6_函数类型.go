package main

import "fmt"

func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return a - b
}
//函数也是一种数据类型,通过type给一个函数类型起名
//Funct她是因各函数类型
type Funct func(int, int) int    //没有函数名字,没有{}

//回调函数,函数有一个参数是函数类型,这个函数就是回调函数
//计算器,可以进行四则运算
//多态,多种形态,调用同一个接口,不同的表现,可以实现不同表现,加减乘除
func Calc(a, b int, fTest Funct) (result int){
	fmt.Println("回调函数:Calc")
	result = fTest(a, b)	//这个函数还没实行
	return
}

func main()  {
	var result int
	result = Add(1, 1)	//传统调用方法
	fmt.Println("传统调用方法:result = ", result)

	//声明一个函数类型的变量,变量名叫fTest
	var fTest Funct
	fTest = Add
	result = fTest(10, 20)	// 等价于Add(10, 20)
	fmt.Println("函数类型调用加法:result2 =", result)

	fTest = Minus
	result = fTest(10, 5)
	fmt.Println("函数类型调用减法:result3 =", result)

	a := Calc(1, 1, Minus)
	fmt.Println("回调函数:a = ", a)
}


//返回结果
/*
传统调用方法:result =  2
函数类型调用加法:result2 = 30
函数类型调用减法:result3 = 5
回调函数:Calc
回调函数:a =  0
 */