package main

import (
	"fmt"
	"time"
)

//主协程退出,其他子协程也跟着退出还没来得及调用
func main() {
	go func() {
		for {
			i := 0
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Println("main i = ", i)
		time.Sleep(time.Second)

		if i == 2 {
			break
		}
	}
}
