package main

import "fmt"

func test(m map[int]string) {
	delete(m, 1)
}

func main() {
	m := map[int]string{1: "make", 2: "yoyo", 3: "go"}
	fmt.Println("m = ", m)
	test(m)
	fmt.Println("m = ", m)
}
