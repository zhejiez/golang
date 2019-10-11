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
	fmt.Printf("SetInfoValue\n")
}

//接受者为指针变量,引用传递
func (p *Person) SetInfoPointer() {
	fmt.Printf("SetInfoPointer\n")
}

func main() {
	//结构体变量是一个指针变量,它能够调用哪些方法,这些方法就是一个集合,简称方法集
	p := &Person{"mike", 'm', 18}
	p.SetInfoPointer()
	//内部做转换,先把指针p,转成*p后再调用
	// (*p).SetInfoValue
	p.SetInfoValue()
}
