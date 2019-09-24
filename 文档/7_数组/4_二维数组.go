package __数组

import (
	"fmt"
)

func main()  {
	//有多少个[]就有多少维,多少循环
	var a [3][4]int

	k := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			k++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d, ", i, j, a[i][j])
		}
		fmt.Printf("\n")
	}
	fmt.Println("a =", a)

	//有三个元素,每个元素又是一维数组[4]int
	b := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("b = ", b)

	//部分初始化,没有初始化的值为0
	c := [3][4]int{{1, 2, 3}, {5, 7}}
	fmt.Println("c = ", c)

	//指定某个元素初始化
	e := [3][4]int{1: {5, 6, 7, 8}}
	fmt.Println("e = ", e)
}