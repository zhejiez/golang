package main

import (
	"fmt"
	"regexp"
)

// 正则表达式查看地址https://studygolang.com/pkgdoc

func main() {
	buf := "abc azc a7c aac 999 a9c t1ac"
	//1)解析规则,他会解析正则表达式,如果成功返回解释器
	//reg1 := regexp.MustCompile(`a.c`)  //匹配a开头,c结尾的字段
	//reg1 := regexp.MustCompile(`a[0-9]c`) //匹配a开头,c结尾的字段,中间为0到9字段
	reg1 := regexp.MustCompile(`azc`)
	if reg1 == nil { //解析失败,返回nil
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg1.FindAllStringSubmatch(buf, -1)
	fmt.Println("restlt1 = ", result1)
}
