package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	list := os.Args //获取命令行参数
	if len(list) != 3 {
		fmt.Println("参数不对")
		return
	}

	srcFileName := list[1]
	dstFileName := list[2]
	if srcFileName == dstFileName {
		fmt.Println("源文件与目标文件名字不能相同")
	}
	//只读方式打开源文件
	sF, err1 := os.Open(srcFileName)
	if err1 != nil {
		fmt.Println("err = ", err1)
		return
	}

	//新建目标文件
	dF, err2 := os.Create(dstFileName)
	if err2 != nil {
		fmt.Println("err2 = ", err2)
		return
	}

	//操作完毕,关闭文件
	defer sF.Close()
	defer dF.Close()

	//核心处理,从源文件读取内容,往目标文件写,读多少写多少
	buf := make([]byte, 4*1024)
	for {
		n, err := sF.Read(buf)
		if err != nil {
			fmt.Println("拷贝完成")
			if err == io.EOF { //文件读取完毕
				break
			}
		}
		dF.Write(buf[:n])
	}
}
