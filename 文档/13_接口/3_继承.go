package main

import (
	"fmt"
)

type Hunmaner interface { //子集
	//方法,只要声明,没有实现,由自定义类型实现
	sayhi()
}

type Personer interface { //超集
	Hunmaner //匿名字段,技能了sayhi
	sing(lrc string)
}

type Student struct {
	name string
	id   int
}

func (tmp *Student) sayhi() {
	fmt.Printf("Student[%s, %d] sayhi\n", tmp.name, tmp.id)
}
func (tmp *Student) sing(lrc string) {
	fmt.Println("Student在看着: ", lrc)
}

func main() {
	//定义一个接口类型的变量
	var i Personer
	s := &Student{"mike", 9}
	i = s

	i.sayhi() // 继承过来的方法
	i.sing("电影")
}
