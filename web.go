package main
import (
	//"database/sql"
	//_ "github.com/mattn/go-sqlite3"
	"fmt"
	"log"
	"net/http"
	//"strconv"
)

func addUser(w http.ResponseWriter, req *http.Request) {
	//userId := req.FormValue("Id")
	name := req.FormValue("name")
	out := name
	log.Println(out)
	fmt.Println(out)
	fmt.Fprintf(w, out)
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