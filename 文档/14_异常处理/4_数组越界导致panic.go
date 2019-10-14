package main

import (
	"fmt"
)

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}
func testb(x int) {
	var a [10]int
	a[x] = 111 //当x为20的时候,导致数组越界产生一个panic
}
func testc() {
	fmt.Println("ccccccccccccccccc")
}
func main() {
	testa()
	testb(20)
	testc()
}
