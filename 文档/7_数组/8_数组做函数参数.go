package __数组

import "fmt"
//数组做函数参数,它是值传递,实参数组每个元素给形参拷贝一分
//形参的数组是实参的复制品
func modify(a [5]int)  {
	a[0] = 666
	fmt.Println("modify a = ", a)
}

func main()  {
	a := [5]int{1,2,3,4,5}
	modify(a)
	fmt.Println("main: a =", a)
}