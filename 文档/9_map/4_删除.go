package main

import "fmt"

func main() {
	m := map[int]string{1: "make", 2: "yoyo", 3: "go"}
	fmt.Println("m = ", m)

	delete(m, 1) //删除key为1的内容
	fmt.Println("m = ", m)
}
