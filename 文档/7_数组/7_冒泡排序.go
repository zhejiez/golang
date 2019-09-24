package __数组

import (
"fmt"
"math/rand"
"time"
)

func main()  {
	//设置种子,只需要一次
	rand.Seed(time.Now().UnixNano())	//以当前系统时间作为种子

	var a [10]int
	n := len(a)

	for i := 0; i < n; i++{
		a[i] = rand.Intn(100)
		fmt.Printf("%d, ", a[i])
	}

	//冒泡排序,挨着的两个元素比较,升序(大于则交换)
	for i := 0; i <n-1; i++ {
		for j :=0; j<n-1-i; j++ {
			if a[j] > a[j+1]{
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
		}
	fmt.Printf("\n排序后:\n")
	for i := 0; i < n; i++{
		fmt.Printf("%d, ", a[i])
	}
	fmt.Printf("\n")
}