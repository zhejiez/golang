package main

import "fmt"

func main()  {
	z := 10
	str := "mike"
	//推导类型匿名函数,没有函数名字,函数定义,没有调用
	f1 := func() {
		fmt.Println("匿名函数无参不调用:z = ", z)
		fmt.Println("匿名函数无参不调用:str = ", str)
	}
	f1()

	//不使用推导类型
	type funct func() //函数没有参数,没有返回值
	var f2 funct
	f2 = f1
	f2()

	//定义匿名函数,无参数调用
	func() {
		fmt.Printf("匿名函数无参数调用: z = %d, str = %s\n", z, str)
	} ()	//后面的()代表调用此匿名函数

	//定义匿名参数,带参数,不调用
	f3 := func(i, j int) {
		fmt.Printf("匿名参数有参数不调用:i = %d, j = %d\n", i, j)
	}
	f3(1, 2)

	//定义匿名函数,带参数调用
	func(i, j int) {
		fmt.Printf("匿名函数有参数调用: i = %d, j = %d\n", i, j)
	} (10, 20)	//后面的()代表调用此匿名函数

	//匿名函数,有参有返回值
	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		} else {
			max = j
			min = i
		}
		return
	}(111, 22)
	fmt.Printf("匿名函数有返回值:x = %d, y = %d\n", x, y)
}


//返回结果
/*
匿名函数无参不调用:z =  10
匿名函数无参不调用:str =  mike
匿名函数无参不调用:z =  10
匿名函数无参不调用:str =  mike
匿名函数无参数调用: z = 0, str = mike
匿名参数有参数不调用:i = 1, j = 2
匿名函数有参数调用: i = 10, j = 20
匿名函数有返回值:x = 111, y = 22
*/