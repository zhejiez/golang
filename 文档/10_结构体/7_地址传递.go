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

func test2(p *Student) {
	p.id = 666
}

func test1(s Student) {
	s.id = 666
	fmt.Println("test01 : ", s)
}

func main() {
	s := Student{1, "mike", 'm', 18, "beiing"}
	test2(&s) //地址传递(引用),形参可以改实参
	fmt.Println("main: ", s)
}
