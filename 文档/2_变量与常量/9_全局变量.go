package main

import "fmt"

func test2()  {
	fmt.Println("test a = ", a)
}
//定义在函数外部的变量是全局变量,在任何地方都可以使用
var a int

func main()  {
	a = 10
	fmt.Println("a = ", a)

	test2()
}

/*
执行结果:
a =  10
test a =  10
*/