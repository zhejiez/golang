package main

import "fmt"

//无参无返回值函数定义
func MyFunc()  {
	fmt.Println("====================================")
	a := 666
	fmt.Println("无参无返回值:a = ", a)
}

//有参无返回值函数定义,普通参数列表,固定参数一定要传参
//定义函数是,在函数名后面的()定义的参数叫型参
//参数传递只能由实参传递给形参
func MyFunc1(a, b int, c string)  {
	fmt.Println("====================================")
	//a = 111
	fmt.Printf("有参无返回值:a = %d, b = %d, c = %s\n", a, b, c)
}

//...int不定参数类型,注意: 不定参数,只能放在形参中的最后一个参数
//func MyFunc2(a ...int, b int)
func MyFunc2(a int, b ...int)  {  //传递的实参可以是一个或者多个
	fmt.Println("====================================")
	fmt.Println("不定参数:a =", a)
	fmt.Println("不定参数:len(b) =", len(b))  //获取用户传递参数的个数
	//for i := 0;i < len(a); i++{
	for _, data := range b{
		fmt.Printf("不定参数:b= %d\n", data)
	}
}

//不定参数的传递
func MyFunc3(tmp ...int)  {
	//全部元素传递给MyFunc2	  MyFunc(b...)
	//MyFunc2(tmp...)
	MyFunc4(tmp[2:]...)		//从tmp[2]开始(包括本身).把后面的所有元素传递过去
	//MyFunc2(tmp[:2]...)		//tmp[0-2](不包括2).元素传递过去
}
func MyFunc4(b ...int)  {		//传递的实参可以是一个或者多个
	fmt.Println("====================================")
	fmt.Println("不定参数传递过来的:len(b) =", len(b))  //获取用户传递参数的个数
	//for i := 0;i < len(a); i++{
	for _, data := range b{
		fmt.Printf("不定参数传递过来的:b= %d\n", data)
	}
}

func main()  {
	//无参无返回函数的调用: 函数名()
	MyFunc()

	//有参无返回值函数调用,函数名(所需参数)
	//调用函数传递的参数叫实参
	MyFunc1(1, 2, "qwe")

	//不定参数的调用,函数名(参数, 参数)
	MyFunc2(1, 2 ,3)
	MyFunc3(1, 2, 3, 4)
}

//返回结果
/*
无参无返回值:a =  666
====================================
有参无返回值:a = 1, b = 2, c = qwe
====================================
不定参数:a = 1
不定参数:len(b) = 2
不定参数:b= 2
不定参数:b= 3
====================================
不定参数传递过来的:len(b) = 2
不定参数传递过来的:b= 3
不定参数传递过来的:b= 4
 */