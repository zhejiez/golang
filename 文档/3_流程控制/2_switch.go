package main

import "fmt"

func main()  {
	//var num int
	//fmt.Printf("请按下楼层:")
	//fmt.Scan(&num)

	//支持一个初始化语句, 初始化语句和变量本身, 以分好分割
	switch num := 4; num {   //switch后面写的是变量本身
	case 1:
		fmt.Println("按下的是1楼")
		break // go语言保留了break关键字,满足条件调出switch判断,默认就包含
		//fallthrough //不跳出switch语句,后面的无条件执行
	case 2:
		fmt.Println("按下的是2楼")
		//fallthrough
	case 3, 4, 5:
		fmt.Println("按下的是yyy楼")
	case 6:
		fmt.Println("按下的是4楼")
	default:
		fmt.Println("按下的是xxx楼")
	}

	score := 92
	switch  { //可以没有条件
	case score > 90:  //case后面可以放条件
		fmt.Println("优秀")
	case score > 80:
		fmt.Println("良好")
	default:
		fmt.Println("一般")
	}
}

//返回结果
/*
按下的是yyy楼
优秀
 */