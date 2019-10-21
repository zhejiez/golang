package main

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

func main() {
	cfg, err := ini.Load("D:\\zxt\\golang\\文档\\15_文本文件处理\\1.conf")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	ves := cfg.Section("").HasKey("app_mode")
	fmt.Println(ves)
}
