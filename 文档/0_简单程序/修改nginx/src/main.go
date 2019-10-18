package main

import (
	"log"
	"net/http"
	"web"
)

func main() {
	http.HandleFunc("/", web.AddUser)
	err := http.ListenAndServe(":8989", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
