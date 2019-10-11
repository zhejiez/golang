package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

//Person类型,实现一个方法
func (tmp *Person) PersonInfo() {
	fmt.Printf("name=%s, sex=%c, age=%d\n", tmp.name, tmp.sex, tmp.age)
}

//有一个学生,继承Person字段,成员和方法都继承
type Student struct {
	Person //匿名字段
	id     int
	addr   string
}

func main() {
	s := Student{Person{"mike", 'n', 17}, 22, "bj"}
	s.PersonInfo()
}
