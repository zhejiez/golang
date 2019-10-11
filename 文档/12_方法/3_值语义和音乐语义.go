package main

import (
	"fmt"
)

type Person struct {
	name string
	sex  byte
	age  int
}

//修改成员变量的值
//参数为普通变量,非指针,值语义,一份拷贝
func (p Person) SetInfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p = ", p)
	fmt.Printf("SetInfoValue &p = %p\n", &p)
}

//接受者为指针变量,引用传递
func (p *Person) SetInfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p = ", p)
	fmt.Printf("SetInfoPointer p = %p\n", p)
}

func main() {
	//var s1 Person
	s1 := Person{"go", 'm', 22}
	fmt.Printf("&s1 = %p\n", &s1) //打印地址

	//值语义
	//s1.SetInfoValue("mike", 'm', 18)
	//fmt.Println("s1 = ", s1)  //打印内容

	//引用语义
	(&s1).SetInfoPointer("mike", 'm', 12)
	fmt.Println("s1 = ", s1)
}
