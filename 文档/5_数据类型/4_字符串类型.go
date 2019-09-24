package main

import "fmt"

func main()  {
	var str1 string  //声明变量
	str1 ="abc"
	fmt.Printf("str1 = %s\n", str1)

	//自动推导类型
	str2 := "mike"
	fmt.Printf("str2 类型是 %T\n", str2)

	//内径函数,len()可以测字符串长度
	fmt.Println("len(star2) = ", len(str2))
}

//返回结果
/*
str1 =  abc
str2 类型是 string
len(star2) =  4
 */