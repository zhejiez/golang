package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person //只有类型,没有名字,匿名字段,继承了Person的成员
	id     int
	addr   string
}

func main() {
	a1 := Student{Person{"make", 'm', 18}, 1, "bj"}
	a1.name = "yoyo"
	a1.sex = 'f'
	a1.age = 22
	a1.id = 666
	a1.addr = "sz"

	a1.Person = Person{"go", 'm', 18}
	fmt.Println(a1.name, a1.sex, a1.age, a1.id, a1.addr)
}
