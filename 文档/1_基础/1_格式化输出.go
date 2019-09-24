package main

import io "fmt" //fmt改为io
import _ "os"	//忽略包

func main()  {
	a := 10
	b := "abc"
	c := 'a'
	d := 3.14
	//%T操作变量所属类型
	io.Printf("%T, %T, %T, %T\n", a, b, c, d)
	// %d:整型格式	%s:字符串格式	%c:字符格式	%f:浮点型
	io.Printf("a = %d, b = %s, c = %c, d = %f\n", a, b, c, d)
	// %v自动匹配格式,不智能
	io.Printf("a = %v, b = %v, c = %v, d = %v\n", a, b, c, d)
}

//返回结果
/*
int, string, int32, float64
a = 10, b = abc, c = a, d = 3.140000
a = 10, b = abc, c = 97, d = 3.14
 */