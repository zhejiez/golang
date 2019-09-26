package main

import "fmt"

func main() {
	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t = ", t)

	//自动推导类型
	t2 := 3.3 + 4.4i
	fmt.Printf("t2 type is %T\n", t2)

	//通过内奸函数,取实部和虚部
	fmt.Println("real(t2)=", real(t2), ", imag(t2)=", imag(t2))
}

//返回结果
/*
t =  (2.1+3.14i)
t2 type is complex128
real(t2)= 3.3 , imag(t2)= 4.4
*/
