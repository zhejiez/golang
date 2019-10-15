package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "3.14 567 agsda 11.23 5. 6.22 kasdn 1.23"

	//解析正则表达式,+匹配前一个字符1次或者多次
	reg := regexp.MustCompile(`\d+\.\d+`)
	if reg == nil {
		fmt.Println("MustCompile err")
		return
	}
	//提取关键信息
	result := reg.FindAllStringSubmatch(buf, -1)
	fmt.Println("result = ", result)
}
