package __数组

import "fmt"

//p指向实现数组a,它是指向数组,他是数组指针
//*p代表指针所指向的内存,就是实参a
func modify1(p *[5]int)  {
	(*p)[0] = 666
	fmt.Println("modify a = ", *p)
}

func main()  {
	a := [5]int{1,2,3,4,5}
	modify1(&a)
	fmt.Println("main: a =", a)
}