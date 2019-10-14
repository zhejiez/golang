package main

import (
	"fmt"
)

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}
func testb(x int) {
	// 设置recover
	defer func() {
		//recover()  //可以打印panic的错误
		//fmt.Println(recover())
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}() //别忘了(),调用此匿名函数
	var a [10]int
	a[x] = 111 //当x为20的时候,导致数组越界产生一个panic
	fmt.Println(a)
}
func testc() {
	fmt.Println("ccccccccccccccccc")
}
func main() {
	testa()
	testb(9)
	testc()
}
