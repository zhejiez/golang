package __切片

import "fmt"

func main()  {
	//自动推导类型
	s1 := []int{1, 2, 3, 4}
	fmt.Println("s1 = ", s1)

	//借助make函数,格式: make(切片类型, 长度, 容量)
	s2 := make([]int, 5, 10)
	fmt.Printf("s2: len = %d, cap = %d\n", len(s2), cap(s2))

	//没有指定容量,与长度一样
	s3 := make([]int, 5)
	fmt.Printf("s3: len = %d, cap = %d\n", len(s3), cap(s3))
}
