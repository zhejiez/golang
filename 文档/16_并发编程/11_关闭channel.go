package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)

	//新建一个goroutine
	go func() {
		for i := 0; i < 7; i++ {
			ch <- i //往通道写数据
		}
		//不需要再写数据时,关闭channel
		//把close(ch)注释掉,程序会一直阻塞在if num, ok := <-ch;ok 那一行
		close(ch)
		//ch <- 1 //关闭后无法发送数据
	}()

	for {
		//如果ok为true,说明管道没有关闭
		if num, ok := <-ch; ok == true {
			fmt.Println("num = ", num)
		} else { //管道关闭
			break
		}
	}
}
