package main

import "fmt"

func test6(x int)  {
	result := 100 / x
	fmt.Println("result = ", result)
}

func main()  {
	//defer延迟调用,main函数结束前调用
	defer fmt.Println("bbbbb")
	defer fmt.Println("aaaaa")
	//调用一个函数导致内存出错
	defer test6(0)
	//后进先出,发生错误仍被执行
	defer fmt.Println("ccccc")

	//defer和匿名函数结合使用
	a := 10
	b := 20
	//defer func() {
	//	fmt.Printf("a = %d, b = %d\n", a, b)
	//}()
	defer func(a, b int) {
		fmt.Printf("a = %d, b = %d\n", a, b)
	}(a, b)		//把参数传递过去,只是没调用相当于(10, 20)
	a = 111
	b = 222
	fmt.Printf("外部:a = %d, b = %d\n", a, b)
}

//返回结果
/*
外部:a = 111, b = 222
a = 10, b = 20
ccccc
aaaaa
bbbbb
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.test6(0x0)
	D:/zxt/awesomeProject/文档/函数/9_defer.go:6 +0xd9
main.main()
	D:/zxt/awesomeProject/文档/函数/9_defer.go:31 +0x262
 */