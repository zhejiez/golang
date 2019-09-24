package main

import "fmt"

func main()  {
	//变量:程序运行期间,可以改变的量,变量声明var
	//常量:程序运行期间,不可以改变的量,常量声明需要const
	const a int = 10
	//a = 20
	fmt.Println("a = ", a)

	const b  = 11.2 //没有使用:=
	fmt.Printf("b type is %T\n", b)
	fmt.Println("b = ", b)
}

//返回结果
/*
a =  10
b type is float64
b =  11.2
 */
