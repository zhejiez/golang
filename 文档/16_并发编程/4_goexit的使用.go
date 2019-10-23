package main

import (
	"fmt"
	"runtime"
)

func test() {
	defer fmt.Println("cccccccccccccc")
	//return //终止此函数
	runtime.Goexit() //终止所在的协程
	fmt.Println("ddddddddddd")
}

func main() {
	//创建一个新的协程
	go func() {
		fmt.Println("aaaaaaaaaa")
		//调用别的函数
		test()
		fmt.Println("bbbbbbbbbb")
	}()

	//特地写一个死循环,不让主协程结束
	for {
	}
}
