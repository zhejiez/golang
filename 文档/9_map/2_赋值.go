package main

import "fmt"

func main() {
	m1 := map[int]string{1: "make", 2: "yoyo"}
	//赋值,如果已经存在key值,修改内容
	fmt.Println("m1 = ", m1)
	m1[1] = "c++"
	fmt.Println("m1 = ", m1)
	m1[3] = "go" //追加,map底层自动扩容,与append类似
	fmt.Println("m1 = ", m1)
}
