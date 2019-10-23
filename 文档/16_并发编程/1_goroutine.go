package main

import (
	"fmt"
	"time"
)

func newTask() {
	for {
		fmt.Println("newTask")
		time.Sleep(time.Second)
	}
}

func main() {
	go newTask() //新建一个协程,新建一个任务

	for {
		fmt.Println("主goroutine")
		time.Sleep(time.Second) //延时1s
	}
}
