//1) go语言以包为管理单位
//2) 每个文件必须先声明包
//3) 程序必须有一个main包
package main

import "fmt"

//入口函数
func main() {      //做大括号必须与函数名同行
	//"hello go"打印到屏幕,Println()会自动换行
	//调用函数,大部分都需要导入包
	/*
		块注释
	 */
	fmt.Println("hello go")
}