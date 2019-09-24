package main

import "fmt"

func main()  {
	s := "王思聪"

	//if到{就是条件,调教同城是关系运算符
	if s == "王思聪"{   //左括号与if在同一行
		fmt.Println("不用工作")
	}
	//if 支持1个初始化语句,初始化语句和判断条件以;分割
	if a := 10; a == 10{   //条件为真,就执行{}语句
		fmt.Println("a == 10" )
	}

	b := 11
	if b == 10 {
		fmt.Println("a == 10")
	} else {  //else后面没有条件
		fmt.Println("b != 10")
	}

	//c := 9
	if c := 9; c == 10{
		fmt.Println("c == 10")
	} else if c < 10{
		fmt.Println("c < 10")
	} else if c > 10{
		fmt.Println("c > 10")
	} else {
		fmt.Println("这是不可能的")
	}
}

//返回结果
/*
不用工作
a == 10
b != 10
c < 10
 */