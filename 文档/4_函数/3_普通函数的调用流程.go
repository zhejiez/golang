package main

import "fmt"

func funca(a int) {
	funcb(a - 1)
	fmt.Println("a = ", a)
}

func funcb(b int) {
	funcc(b - 1)
	fmt.Println("b = ", b)
	return
}

func funcc(c int)  {
	fmt.Println("c = ", c)
}

func main()  {
	funca(3)  //函数调用
	fmt.Println("main")
}

//返回结果
/*
c =  1
b =  2
a =  3
main
 */