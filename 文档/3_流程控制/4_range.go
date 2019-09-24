package main

import "fmt"

func main()  {
	str := "abc"

	//通过for打印每个字符
	for i:= 0; i < len(str); i++{
		fmt.Printf("str[%d],=%c\n", i, str[i])
	}

	//迭代打印每个元素,默认返回2个值,一个元素位置,一个元素本身
	for i, data := range str {
		fmt.Printf("str[%d],=%c\n", i, data)
	}

	for i:= range str { //第二个返回值,默认丢弃,返回元素的位置(下标)
		fmt.Printf("str[%d],=%c\n", i, str[i])
	}
}

//返回结果
/*
str[0],=a
str[1],=b
str[2],=c
str[0],=a
str[1],=b
str[2],=c
str[0],=a
str[1],=b
str[2],=c
 */