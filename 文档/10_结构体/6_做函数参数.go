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

func test01(s Student) {
	s.id = 666
	fmt.Println("test01 : ", s)
}

func main() {
	s := Student{1, "mike", 'm', 18, "beiing"}
	test01(s) //值传递,形参无法改实参
	fmt.Println("main: ", s)
}
