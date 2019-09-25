package main

import (
	"fmt"
)

type Student struct {
	id   int
	name string
	sex  byte
	age  int
	addr string
}

func main() {
	//1,指针有合法指向,才操作成员
	//先定义一个普通结构体变量
	var s Student
	//在定义一个指针变量,保存s的地址
	var p1 *Student
	p1 = &s

	//通过指针操作人员,p1.id和(*p1).id完全等级,只能使用.运算符
	p1.id = 1
	(*p1).name = "mike"
	p1.sex = 'm'
	p1.age = 18
	p1.addr = "beijing"
	fmt.Println("p1 = ", *p1)

	//2,通过new申请一个结构体
	p2 := new(Student)
	p2.id = 1
	(*p2).name = "mike"
	p2.sex = 'm'
	p2.age = 18
	p2.addr = "beijing"
	fmt.Println("p2 = ", *p2)
}
