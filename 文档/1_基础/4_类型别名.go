package main

import "fmt"

func main()  {
	//给int64起一个别名bigint
	type bigint int64

	var a bigint  //等价于var a int64
	fmt.Printf("a type is %T\n", a )

	//多个别名
	type (
		long int64
		char byte
	)
	var b long = 11
	var ch char = 'a'
	fmt.Printf("b = %d, ch = %c", b, ch)
}
//返回结果
/*
a type is main.bigint
b = 11, ch = a
 */