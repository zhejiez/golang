package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Stdout.Close()		//关闭后无法输出
	//fmt.Println("a y k?") //往标准输出设备(屏幕)写内容

	//标准设备文件,默认已打开,用户可直接使用
	os.Stdout.WriteString("a y k?\n")

	os.Stdin.Close() //关闭后无法输入
	var a int
	fmt.Println("请输入a")
	fmt.Scan(&a) //从标准输入设备住读取内容,放在a中
	fmt.Println("a = ", a)
}
