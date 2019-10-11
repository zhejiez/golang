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

//Student也实现了一个方法,这个方法和Person方法同名,这种方法叫重写
func (tmp *Student) PersonInfo() {
	fmt.Println("Student: tmp = ", tmp)
}

func main() {
	s := Student{Person{"mike", 'n', 17}, 22, "bj"}
	//就近原则
	s.PersonInfo() //到底调用Person,还是Student,结果调用Student

	//显示调用继承的方法
	s.Person.PersonInfo()
}
