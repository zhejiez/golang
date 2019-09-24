package main

import "fmt"

func main()  {
	var a int = 10
	//每个变量有两层含义:变量的内存,变量的地址
	fmt.Printf("变量的内存%d\n", a)
	fmt.Printf("变量的地址%v\n", &a)

	//保存某个变量的地址,需要指针类型		*int保存int的地址
	//声明(定义).定义只是特殊声明
	//定义变量p,类型为*int
	var p *int
	//fmt.Println("p = ", p)
	//*p = 99	//err,因为p没有合法指向

	p = &a //指针变量指向谁,就把谁的地址复制给指针变量
	fmt.Printf("p = %v, &a - %v\n", p, &a)

	*p = 666 //*p操作的不是p的内存,是p所指向的内存(就是a)
	fmt.Printf("*p = %v, a = %v\n", *p, a)
}