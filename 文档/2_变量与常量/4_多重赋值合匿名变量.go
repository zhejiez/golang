package main

import "fmt"

//go函数可以返回多个值
func test() (a, b, c int){
	return 1, 2,3
}
func main()  {
	a, b := 10, 20
	//交换两个变量的值
	var tmp int
	tmp = a
	a = b
	b = tmp
	fmt.Printf("a = %d, b = %d\n", a, b)

	i := 10
	j := 20
	i, j = j, i
	fmt.Printf("i = %d, j = %d\n", i, j)

	i = 11
	j = 22
	//匿名变量,丢弃数据不处理,_匿名变量配合函数返回值才有优势
	tmp, _ = i, j
	fmt.Println("tmp = ", tmp)

	var c, d, e int
	c, d, e = test() //return 1, 2, 3
	fmt.Printf("c = %d, d = %d, e = %d\n", c, d, e)
	//使用单个变量
	_, d, e = test()//return 1, 2, 3
	fmt.Printf("d = %d, e = %d\n", d, e)
 }

//返回结果
/*
i = 20, j = 10
tmp =  11
c = 1, d = 2, e = 3
d = 2, e = 3
 */