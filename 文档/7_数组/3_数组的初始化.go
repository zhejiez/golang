package __数组

import "fmt"

func main()  {
	//声明定义同时赋值,叫初始号
	//1, 全部初始号
	var a [5]int = [5]int{1,2,3,4,5}
	fmt.Println("a = ", a)

	b := [5]int{1,2,3,4,5}
	fmt.Println("b = ", b)

	//2,部分初始号,没有初始化元素,自动赋值为0
	c := [5]int{1,2,3}
	fmt.Println("c = ", c)

	//指定某个元素初始化
	d := [5]int{2: 10,4 : 20}
	fmt.Println("d = ", d)
}