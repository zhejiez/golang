package main

import (
	"fmt"
	"strings"
)

func main() {
	//Contains:查看hellogo是否包含go,是就返回true,不包含就返回false
	fmt.Println(strings.Contains("Contains: hellogo", "go"))
	fmt.Println(strings.Contains("Contains: hellogo", "a"))

	//Joins 组合
	s := []string{"abc", "hello", "mike", "go"}
	buf := strings.Join(s, "#")
	fmt.Println("Joins: buf = ", buf)

	//Index,查找子串位置位置,不存在为-1
	fmt.Println(strings.Index("Index: abchello", "hello"))
	fmt.Println(strings.Index("Index: abchello", "g"))

	buf = strings.Repeat("Index: go", 3)
	fmt.Println("Index: buf = ", buf) //gogogo

	//Split 以指定的分隔符拆分
	buf = "hello@abc@go@mike"
	s2 := strings.Split(buf, "@")
	fmt.Println("Split: s2 = ", s2)

	// Trim 去掉两头的字符
	buf = strings.Trim("    asd        ", "     ")
	fmt.Printf("trim: buf = #%s#\n", buf)

	//去掉空格,八元素放入切片中
	s3 := strings.Fields("    asd q w        ")
	//fmt.Printf("s3 = #%s#\n", s3)
	for i, data := range s3 {
		fmt.Println(i, ", ", data)
	}
}
