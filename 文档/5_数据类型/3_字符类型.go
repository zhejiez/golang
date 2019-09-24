package main

import "fmt"

func main()  {
	var ch byte //声明字符串类型
	ch = 98
	//%c以字符方式打印
	fmt.Printf("%c, %d\n", ch, ch)

	ch = 'a' //字符,单引号
	fmt.Printf("%c, %d\n", ch, ch)

	//大小转小写,小写转大写
	fmt.Printf("大写: %d, 小写: %d\n", 'A', 'a')
	fmt.Printf("大写转小写: %c\n", 'A' + 32)
	fmt.Printf("小写转大写: %c\n", 'a' - 32)

	// '\'以反写光开头的字符是转移字符,不会输出到屏幕
	fmt.Printf("hello go%c", '\n')
	fmt.Printf("hello itcast")
}

//返回结果
/*
b, 98
a, 97
大写: 65, 小写: 97
大写转小写: a
小写转大写: A
hello go
hello itcast
 */