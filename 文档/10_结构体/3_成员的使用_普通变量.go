package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//顺序初始化,每个成员必须初始化
	var s Student

	//操作成员,使用(.)运算符
	s.id = 1
	s.name = "mike"
	s.sex = 'm'
	s.age = 18
	s.addr = "beijing"
	fmt.Println("s = ", s)
}
