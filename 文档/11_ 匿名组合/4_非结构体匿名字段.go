package main

import (
	"fmt"
)

type mystr string //自定义类型,给一个类型改名
type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person     //结构体匿名字段
	int    int //基础类型的匿名
	mystr
}

func main() {
	s := Student{Person{"mike", 'm', 18}, 666, "he"}
	fmt.Printf("s = %+v\n", s)
	fmt.Println(s.name, s.age, s.sex, s.int, s.mystr)
	fmt.Println(s.Person, s.int, s.mystr)
}
