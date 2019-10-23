package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)
	fmt.Println("len(ch) = %d, cap(ch)= %d\n", len(ch), cap(ch))

	//新建协程
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i //网chan写内容
			fmt.Printf("子协程[%d]: len(ch) = %d, cap(ch)= %d\n", i, len(ch), cap(ch))
		}
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		num := <-ch //读取管道中内容,没有内容阻塞
		fmt.Println("num = ", num)
	}
}
