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
	Person //只有类型,没有名字,匿名字段,继承了Person的成员
	id     int
	addr   string
	name   string
}

func main() {
	var s Student

	//默认规则(就近原则),如果能在本作用域找到此成员,就操作此成员,没找到就找继承的字段
	s.name = "mike" //操作的是Person还是Student的name?
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"
	//显示调用
	s.Person.name = "yoyo"
	fmt.Printf("s = %+v\n", s)
}
