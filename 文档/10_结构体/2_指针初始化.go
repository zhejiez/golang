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
	//顺序初始化,每个成员必须初始化,别忘了&
	var p1 *Student = &Student{1, "mike", 'm', 18, "beijing"}
	fmt.Println("p1 = ", *p1)

	//指定成员初始化,没有初始化的成员,自动赋值为0
	p2 := &Student{name: "mike", addr: "beijing"}
	fmt.Printf("p2 type is %T\n", p2)
	fmt.Println("p2 = ", *p2)
}
