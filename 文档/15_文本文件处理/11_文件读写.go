package main

import (
	"fmt"
	"io"
	"os"
)

func WriteFile(path string) {
	var buf string
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	} else {
		for i := 0; i < 10; i++ {
			// "i = %d\n",这个字符串存储在buf中
			buf = fmt.Sprintf("i = %d\n", i)
			//n , err := f.WriteString(buf)    //用来排错
			//if err != nil {
			//	fmt.Println("err = ", err)
			//}
			//fmt.Println("n = ", n)
			f.WriteString(buf)
		}
	}
	//使用完毕,关闭文件
	defer f.Close()
}

func ReadFile(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()
	buf := make([]byte, 1024*2)

	//n代表从文件读取内容长度
	n, err1 := f.Read(buf)
	if err1 != nil && err1 != io.EOF { //读取文件出错,同事没到结尾
		fmt.Println("err1 = ", err1)
		return
	}
	fmt.Println("buf = ", string(buf[:n]))
}

func main() {
	path := "D:\\zxt\\golang\\文档\\0_简单程序\\h.txt"
	WriteFile(path)
	ReadFile(path)
}
