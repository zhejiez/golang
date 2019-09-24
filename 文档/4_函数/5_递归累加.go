package main

import "fmt"

//实现1+2+3+......100
func test1() (sum int) {
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return
}

func test2(i int) int {
	if i == 1 {
		return 1
	}
	return i + test2(i-1)
}

func test3(i int) int {
	if i == 100 {
		return 100
	}
	return i + test3(i+1)
}

func main()  {
	var sum int
	//sum = test1()
	//sum = test2(5)
	sum = test3(1)
	fmt.Println("sum =", sum)
}
