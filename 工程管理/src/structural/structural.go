package structural

import (
	"fmt"
)

type stu struct {
	id int
}

type Stu struct {
	//id int  //如果首字母小写,只能这个包用
	Id int
}

func myFunc() {
	fmt.Println("this is myFunc")
}

func MyFunc() {
	fmt.Println("结构体")
}
