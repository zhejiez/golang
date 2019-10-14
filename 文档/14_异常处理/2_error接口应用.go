package main

import (
	"errors"
	"fmt"
)

func MyDia(a, b int) (result int, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = a / b
	}
	return
}
func main() {
	result, err := MyDia(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("reslut = ", result)
	}

}
