package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个定时器,设置时间为2s,2s后往tim通道写内容(当前时间)
	timer := time.NewTimer(2 * time.Second)

	////验证time.NewTimer()时间到了,只会响应一次
	//for {
	//	<-timer.C
	//	fmt.Println("时间到")
	//}

	fmt.Println("当前时间:", time.Now())
	//2s后往timer.C写数据,有数据后,就可以读取
	t := <-timer.C //channel没有数据前后阻塞
	fmt.Println("t = ", t)
}
