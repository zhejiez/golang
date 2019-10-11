package main

import (
	//"database/sql"
	//_ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
	"net/http"
	//"strconv"
)

//http://127.0.0.1:8989/user?Id=1&id=o&name=guo
func addUser(w http.ResponseWriter, req *http.Request) {
	//userId := req.FormValue("Id")
	id, name := req.FormValue("id"), req.FormValue("name")
	fmt.Println(id)
	fmt.Println(name)
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world!")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/user", addUser)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
