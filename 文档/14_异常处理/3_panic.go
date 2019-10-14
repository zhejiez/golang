package main

import (
	"fmt"
)

func testa() {
	fmt.Println("aaaaaaaaaaaaaaaaa")
}
func testb() {
	fmt.Println("bbbbbbbbbbbbbbbbb")
	panic("this is a panic test")
}
func testc() {
	fmt.Println("ccccccccccccccccc")
}
func main() {
	testa()
	testb()
	testc()
}
