package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(3 * time.Second)
	//timer.Reset(1*time.Second)  //重新设置时间,1秒后执行
	go func() {
		<-timer.C
		fmt.Println("此协程可以打印了,因为定时器时间到")
	}()
	//timer.Stop()  //停止定时器
	for {

	}
}
