package main

import "fmt"

func main()  {
	var a int
	fmt.Printf("请输入变量a: ")
	//堵塞等待用户的输入
	//fmt.Scanf("%d, &a", &a)   //别忘了&
	fmt.Scan(&a)
	fmt.Println("a = ", a)
}

//返回结果
/*
请输入变量a: 123
a =  123
 */