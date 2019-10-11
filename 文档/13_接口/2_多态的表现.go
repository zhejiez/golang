package main

import (
	"fmt"
)

//定义接口类型
type Hunmaner interface {
	//方法,只要声明,没有实现,由自定义类型实现
	sayhi()
}

type Student struct {
	name string
	id   int
}

func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}

type Teacher struct {
	addr  string
	group string
}

//Teacher实现了此方法
func (tmp *Teacher) sayhi() {
	fmt.Printf("Teacher[%s, %s] sayhi\n", tmp.addr, tmp.group)
}

type MyStr string

//MyStr实现了此方法
func (tmp *MyStr) sayhi() {
	fmt.Printf("MyStr[%s] sayhi\n", *tmp)
}

//定义一个普通函数,函数的参数为接口类型
//只有一个函数,可以有不同表现
func WhoSayHi(i Hunmaner) {
	i.sayhi()
}
func main() {
	s := &Student{"mike", 777}
	t := &Teacher{"bj", "go"}
	var str MyStr = "hello mike"

	//调用同一个函数,不同表现,多态
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	x := make([]Hunmaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str
	//第一个返回下标,第二个返回下标所对应的值
	for _, i := range x {
		i.sayhi()
	}
}
