package main

import "fmt"

func main()  {
	//变量,程序运行期间,可以改变的量
	//1, 生命格式  var变量名 类型
	//2, 只是声明没有初始化,默认值为0
	//3, 同一{}里,声明变量名是唯一的
	var a int  //变量必须使用
	//4, 可以同时声明多个变量
	//var b, c int
	fmt.Println("a = ", a)
	a = 3  //变量的赋值
	fmt.Println("a = ", a)
	//变量的初始化,声明变量,同时赋值
	var b int = 3
	b = 20
	fmt.Println("b = ", b)
	//自动推导类型,必须舒适化,通过初始化的值确定类型(常用
	c := 30
	//%T打印变量所属类型
	fmt.Printf("c type is %T\n", c)
}

//返回结果
/*
a =  0
a =  3
b =  20
c type is int
 */