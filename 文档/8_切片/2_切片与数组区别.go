package __切片

import "fmt"

func main()  {
	//切片与数组的区别
	//数组[]里面的长度是固定的一个常量,数组不能修改长度,len和cap永远是5
	a := [5]int{}
	fmt.Printf("a: len = %d, cap = %d\n", len(a), cap(a))

	//切片,[]里面为空,或者...,切片的长度或容量可以不固定
	s := []int{}
	fmt.Printf("s: len = %d, cap = %d\n", len(s), cap(s))

	s = append(s, 11)  //给切片末尾垂加一个成员
	fmt.Printf("append: len = %d, cap = %d\n", len(s), cap(s))
}