package main

import (
	"bufio"
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

func ReadFileLine(path string) {
	//打开文件
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	defer f.Close()

	//新建一个缓冲区,把内容先放在缓冲区
	r := bufio.NewReader(f)
	for {
		//遇到'\n'结束读取,但是'\n'也读取进入
		buf, err := r.ReadBytes('\n')
		if err != nil {
			if err == io.EOF { //文件已经结束
				break
			}
			fmt.Println("err = ", err)
		}
		fmt.Println(err)
		fmt.Printf("buf = #%s#\n", string(buf))
	}
}

func main() {
	path := "D:\\zxt\\golang\\文档\\0_简单程序\\h.txt"
	WriteFile(path)
	ReadFileLine(path)
}
