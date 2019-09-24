package main

import "fmt"

func test4() int {
	//函数被调用时,x才分配空间,才初始为0
	var x int //没有初始化,值为0
	x++
	return x * x	//函数调用完毕,x自动释放
}

//函数的返回值是一个匿名函数,返回一个函数类型
func test5() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main()  {
	a := 10
	str := "mike"

	func() {
		//闭包以引用方式捕获外部变量
		a = 666
		str = "go"
		fmt.Printf("内部: a = %d, str = %s\n", a, str)
	}()  //()代表直接调用
	fmt.Printf("外部:a = %d, str = %s\n", a, str)

	//返回值为一个匿名函数,返回一个函数类型,通过f来调用闭包函数
	//它不关心这些捕获了的变量和常量是否超出作用域,所以只有闭包还在使用它,这些变量就一直存在
	fmt.Println("传统调用: ", test4())
	fmt.Println("传统调用: ", test4())
	f := test5()
	fmt.Println("闭包调用: ", f())
	fmt.Println("闭包调用: ", f())
	fmt.Println("闭包调用: ", f())
}
//返回结果
/*
内部: a = 666, str = go
外部:a = 666, str = go
传统调用:  1
传统调用:  1
闭包调用:  1
闭包调用:  4
闭包调用:  9
 */