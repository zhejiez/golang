package main

import (
	"fmt"
)

//定义接口类型
type Hunmaner interface {
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

func (tmp *Student) sing(lrc string) {
	panic("implement me")
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

func main() {
	//超集可以转换为子集,反过来不可以
	var iPro Personer
	var i Hunmaner
	iPro = &Student{"mike", 666}
	//iPro = i  //error
	i = iPro //超集可以转换为子集
	i.sayhi()
}
