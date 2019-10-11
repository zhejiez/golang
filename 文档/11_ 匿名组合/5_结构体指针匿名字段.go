package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	*Person     //结构体匿名字段
	id      int //基础类型的匿名
	addr    string
}

func main() {
	s1 := Student{&Person{"mike", 'm', 16}, 666, "bj"}
	//fmt.Println("s1 = ", s1)
	fmt.Println(s1.Person, s1.id, s1.addr)

	var s2 Student
	s2.Person = new(Person) //分配空间
	s2.name = "yoyo"
	s2.sex = 'm'
	s2.age = 18
	s2.id = 222
	s2.addr = "hz"
	println(s2.name, s2.sex, s2.age, s2.id, s2.addr)
}
