package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

//待遇接受者的函数叫方法
//只要接受者类型不一样,这个方法就算同名,也是不同的方法,不会出现重复定义函数的错误
func (tmp Person) PrintInfo() {
	fmt.Println("tmp = ", tmp)
}

//type Person *int
// Person为接受者类型，本身不能是指针
//通过一个函数,给成员赋值
func (p *Person) SetInfo(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
}
func main() {
	//定义同事初始化
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	//定义一个结构体边框
	var p2 Person
	(&p2).SetInfo("yoyo", 'f', 22)
	p2.PrintInfo()
}
