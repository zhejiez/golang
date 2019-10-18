package main

import (
	"encoding/json"
	"fmt"
)

/*
{
	"Company": "itcast",
	"Subjects": ["Go", "C++", "Python", "test"],
	"IsOk": true,
	"Price": 555.32
}
*/
type IT struct {
	//Company string
	Company  string   `json:"company"` //变成小写,二次编码
	Subjects []string `json:"-"`       //此字段不会输出
	IsOk     bool
	Price    float64
}

func main() {
	//定义一个结构体变量,同事初始化
	s := IT{"itcast", []string{"Go", "C++", "Python", "test"}, true, 555.32}
	//编码,根据内容生成json文本
	//{"Company":"itcast","Subjects":["Go","C++","Python","test"],"IsOk":true,"Price":555.32}
	//buf, err := json.Marshal(s)  /都在一行
	buf, err := json.MarshalIndent(s, "", " ") //分行显示更清晰
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("buf = ", string(buf))
}
