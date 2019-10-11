package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

func (p Person) SetInfoValue() {
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

//接受者为指针变量,引用传递
func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer%p, %v\n", &p, p)
}

func main() {
	p := Person{"mike", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)

	//方法值 f :=SetInfoPointer //隐藏接受者
	//方法表达式
	f := (*Person).SetInfoPointer
	f(&p) //显示把接受者传递过去 -----> p.SetInfoPointer

	f2 := (Person).SetInfoValue
	f2(p) //显示把接受者传递过去 -----> p.SetInfoValue
}
