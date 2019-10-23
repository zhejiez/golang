package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个无缓存channel
	ch := make(chan int, 0)
	//len(ch)缓冲区剩余数据个数,cap(ch)缓冲区大小
	fmt.Printf("len(ch)= %d, cap(ch)= %d\n", len(ch), cap(ch))
	//新建协程
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("子协程: i=%d\n", i)
			ch <- i //网chan写内容
			//fmt.Printf("len(ch) = %d, cap(ch)= %d \n", len(ch), cap(ch))
		}
	}()
	time.Sleep(2 * time.Second)

	for i := 0; i < 3; i++ {
		num := <-ch //读取管道中内容,没有内容阻塞
		fmt.Println("num = ", num)
	}
}
