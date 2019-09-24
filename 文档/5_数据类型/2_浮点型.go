package main

import "fmt"

func main()  {
	var f1 float32
	f1 = 3.13
	fmt.Println("f1 = ", f1)
	//自动推导类型
	f2 := 3.33
	fmt.Printf("f2 type is %T\n", f2) //默认为float64
	//float64存储小数比float32更准确
}

//返回结果
/*
f1 =  3.13
f2 type is float64
 */