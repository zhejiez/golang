package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	//Company string
	Company  string   `json:"company"`  //变成小写,二次编码
	Subjects []string `json:"subjects"` //此字段不会输出
	IsOk     bool
	Price    float64
}

func main() {
	jsonBuf := `
{
	"company": "itcast",
	"subjects": ["Go", "C++", "Python", "test"],
	"IsOk": true,
	"Price": 555.32
}`
	var tmp IT                                   //定义一个结构体变量
	err := json.Unmarshal([]byte(jsonBuf), &tmp) //第二个参数(tmp)要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp = %+v\n", tmp)

	type IT2 struct {
		Subjects []string `json:"subjects"` //二次编码
	}
	var tmp2 IT2
	err = json.Unmarshal([]byte(jsonBuf), &tmp2) //第二个参数(tmp)要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("tmp2 = %+v\n", tmp2)
}
