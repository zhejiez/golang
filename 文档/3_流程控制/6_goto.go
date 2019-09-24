package main

import "fmt"

func main()  {
	//break  break is not in a loop, switch, or select
	//continue  continue is not in a loop

	//goto可以用在任何地方,但是不可跨函数使用
	fmt.Println("11111111111")
	goto End //goto是关键字, End是用户气的名字,他叫标签
	fmt.Println("22222222222")

End:  //使用goto标签
	fmt.Println("33333333333")
}