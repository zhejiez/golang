package __切片

import "fmt"

func main()  {
	srcSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6}

	copy(dstSlice, srcSlice)
	fmt.Println("dst = ", dstSlice)
}
