package main

import "fmt"

var a byte //全局变量

func main()  {
	var a int //局部变量

	//1,不同作用域,允许定义同名变量
	//2,使用变量原装,就近原装
	fmt.Printf("1: %T\n", a)	//int

	{
		var a float64
		fmt.Printf("2: %T\n", a)	//float64
	}
	test3()
}

func test3()  {
	fmt.Printf("3: %T\n", a)	//uint8
}

/*
执行结果:
1: int
2: float64
3: uint8
*/