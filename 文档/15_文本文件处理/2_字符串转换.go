package main

import (
	"fmt"
	"strconv"
)

func main() {
	//转换为字符串后追加到字节数组
	slice := make([]byte, 0, 1024)
	slice = strconv.AppendBool(slice, true)
	//第二个数为要追加的数,第三个位10禁止方式追加
	strconv.AppendInt(slice, 1234, 10)
	slice = strconv.AppendQuote(slice, "abcasd")

	fmt.Println("slice = ", string(slice))
}
