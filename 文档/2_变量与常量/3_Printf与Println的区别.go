package main

import "fmt"

func main()  {
	a := 10
	//Println : 一段一段处理,自动换行
	fmt.Println("a = ", a)
	//Printf  : 格式化输出,把a的内容放在%d的位置
	//"a = 10\n" 这个字符串输出到屏幕,"\n"代表换行符
	fmt.Printf("a = %d\n" , a)
	b := 20
	c := 30
	fmt.Println("a = ", a, "b = ", b, "c = ", c)
	fmt.Printf("a = %d, b = %d, c = %d\n", a, b, c)
}

//返回结果
/*
a =  10
a = 10
a =  10 b =  20 c =  30
a = 10, b = 20, c = 30
 */