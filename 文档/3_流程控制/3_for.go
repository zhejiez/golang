package main

import "fmt"

func main()  {
	// for  出事条件 ;  条件判断 ;  条件变化{}
	//1+2+3 ...... 100 雷军
	sum :=0
	//1)初始化条件 i := 1
	//)2判断条件是否为真, i <= 100, 如果为真,执行循环体,如果为假,调出循环
	//3)条件变化i++
	//4)重复2,3,4
	for i:= 1; i<= 100; i++{
		sum = sum + i
	}
	fmt.Println("sum = ",sum)
}

//返回结果
/*
sum =  5050
 */