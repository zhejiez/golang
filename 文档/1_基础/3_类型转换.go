package main

import "fmt"

func main()  {
	//这种不能转换的类型,叫做不兼容类型
	var flag bool
	flag = true
	fmt.Printf("flag = %t\n", flag)
	/*
	bool类型不能转换为int
	fmt.Printf("flag = %d\n", int(flag))
	0就是假,1就是真,整形也不能转换为bool
	flag = bool(1)
	*/

	var ch byte
	ch = 'a' // 字符型本质上是整形
	var t int
	t = int(ch)  //类型转换,吧ch的值取出来后,转成int再给t赋值
	fmt.Println("t = ", t)
}

//返回结果
/*
flag = true
t =  97
 */