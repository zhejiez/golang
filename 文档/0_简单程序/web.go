package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//http://127.0.0.1:8989/user?Id=1&id=o&name=guo
func addUser(w http.ResponseWriter, req *http.Request) {
	userId := req.FormValue("Id")
	//id, name := req.FormValue("id"), req.FormValue("name")
	nginxconf(name)
}

func nginxconf(name string) {
	path := "D:\\zxt\\golang\\文档\\0_简单程序\\1.txt"
	f, err := os.Create(path)
	var buf string
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//写入文件
	fmt.Sprintf("if  ($request_uri  ~* \"%s\") {\nreturn 403;\n}\n", name)
	f.Write()
	//使用完毕,关闭文件
	defer f.Close()
}

func main() {
	http.HandleFunc("/", addUser)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
