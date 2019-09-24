package calc

import "fmt"
//优先执行init函数,一般配合: _ "fmt'使用
func init()  {
	fmt.Println("calc init函数")
}
//func add(a, b int) int {
func Add(a, b int) int {	//不同目录函数名必须大写
	return a + b
}

func Minus(a, b int) int {
	return a - b
}