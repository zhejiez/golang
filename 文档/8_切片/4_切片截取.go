package __切片

import "fmt"

func main()  {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// [low:high:max] 取下标从low开始的元素, len=high-low, cap=max-low
	s1 := array[:]	//[0:len(array):len(array)]不指定容量和长度一样
	fmt.Println("s1 = ", s1)
	fmt.Printf("s1: len = %d, cap = %d\n", len(s1), cap(s1))

	//操作某个元素与操作数组的方式一样
	data := array[1]
	fmt.Println("data = ", data)

	s2 := array[3:6:7]	//a[3], a[4], a[5] 	len = 6-3  cap = 7-3
	fmt.Println("s2 = ", s2)
	fmt.Printf("s2: len = %d, cap = %d\n", len(s2), cap(s2))

	s3 := array[:6]			//常用
	fmt.Println("s3 = ", s3)
	fmt.Printf("s3: len = %d, cap = %d\n", len(s3), cap(s3))

	s4 := array[3:]  //从下标为3开始,到结尾
	fmt.Println("s4 = ", s4)
	fmt.Printf("s4: len = %d, cap = %d\n", len(s4), cap(s4))
}