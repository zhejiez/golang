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
	p.SetInfoPointer() //传统调用方式

	//保存方式入口地址
	pFunc := p.SetInfoPointer //这个就是方法值,调用函数时无需再传递接受者,隐藏了接受者
	pFunc()                   //等价于p.SetInfoPointer()

	vFunc := p.SetInfoValue
	vFunc()
}
