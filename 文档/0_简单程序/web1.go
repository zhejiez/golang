package main

import (
	"fmt"
	"log"
	"net/http"
	"web"
)

func AddUser(w http.ResponseWriter, req *http.Request) {
	//配置文件目录
	//var path string = ("D:\\zxt\\golang\\文档\\0_简单程序\\")
	id, name := req.FormValue("id"), req.FormValue("name")
	if id == "h" {
		fmt.Println(id, name)
	} else {
		fmt.Println("错误的子域名:", id)
	}
}

func main() {
	http.HandleFunc("/", web.AddUser)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
