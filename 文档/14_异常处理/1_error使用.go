package main

import (
	"errors"
	"fmt"
)

func main() {
	//var err error = fmt.Errorf("%s", "this in normol err")  //不推荐
	err := fmt.Errorf("%s", "this is normol err")
	fmt.Println("err = ", err)

	err2 := errors.New("this is normal err2")
	fmt.Println("err2 = ", err2)
}
