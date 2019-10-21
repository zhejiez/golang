package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个数为要追加的数,第三个位10禁止方式追加
	strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcasd")

	fmt.Println("slice = ", string(slice)) //转换为string打印

	//其他类型转换为字符串
	var str string
	str = strconv.FormatBool(false)
	//'f'指打印格式,以小幅方式,-1指小数点位数(紧缩模式),64以float64处理
	str = strconv.FormatFloat(3.14, 'f', -1, 32)

	//整形转字符串,常用
	str = strconv.Itoa(6666)

	//字符串转其他类型
	var flag bool
	var err error
	flag, err = strconv.ParseBool("t1rue")
	if err == nil {
		fmt.Println("flag = ", flag)
	} else {
		fmt.Println("err = ", err)
	}
	fmt.Println("str = ", str)

	//把字符串转换为整形,常用
	a, _ := strconv.Atoi("123")
	fmt.Println("a = ", a)
}
