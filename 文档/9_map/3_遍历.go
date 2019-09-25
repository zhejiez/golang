package main

import (
	"fmt"
)

func main() {
	m := map[int]string{1: "make", 2: "yoyo", 3: "go"}
	for key, value := range m {
		fmt.Printf("%d ------> %s\n", key, value)
	}
	valie, ok := m[1]
	if ok == true {
		fmt.Println("m[1] = ", valie)
	} else {
		fmt.Println("key不存在")
	}
}
