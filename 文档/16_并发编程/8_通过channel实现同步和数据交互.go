package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	defer fmt.Println("主协程也结束")
	go func() {
		defer fmt.Println("子协程调用完毕")
		for i := 2; i < 2; i++ {
			fmt.Println("子协程: i = ", i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协程,工作完毕"
	}()
	str := <-ch //没有数据之前阻塞
	fmt.Println("str = ", str)
}
