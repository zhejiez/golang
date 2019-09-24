package main

import (
	"fmt"
)

//无参一个返回值,有返回值的函数需要通过return返回
func mfume() (result int) {
	result = 666
	return
}

//多个返回值
func mfume2() (a, b, c int) {
	a, b, c = 1, 2, 3
	return
}

func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	}else {
		max = b
		min = a
	}
	return
}
func main()  {
	//无参有返回值函数调用
	b := mfume()
	fmt.Println("单个返回值:b =", b)

	//无参有返回值函数调用
	a, b, c := mfume2()
	fmt.Printf("多个返回值:a = %d, b = %d, c = %d\n", a, b, c)

	//有参有返回值的使用
	max, min := MaxAndMin(10, 20)
	fmt.Printf("有参有返回值:max = %d, min = %d\n", max, min)
	//通过匿名变量丢弃某个返回值
	a, _ = MaxAndMin(10, 20)
	fmt.Printf("丢弃某个返回值:a = %d\n", a)
}

//返回结果
/*
单个返回值:b = 666
多个返回值:a = 1, b = 2, c = 3
有参有返回值:max = 20, min = 10
丢弃某个返回值:a = 20
 */