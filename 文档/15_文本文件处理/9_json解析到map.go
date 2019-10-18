package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	jsonBuf := `
{
	"company": "itcast",
	"subjects": ["Go", "C++", "Python", "test"],
	"IsOk": true,
	"Price": 555.32
}`
	//创建一个map
	m := make(map[string]interface{}, 4)

	err := json.Unmarshal([]byte(jsonBuf), &m) //第二个参数(tmp)要地址传递
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Printf("m = %+v\n", m)

	//var str string
	//str = string(m["company"])  //无法转换
	//fmt.Println("str = ", str)

	//类型断言, 值是value类型
	var str string
	for key, value := range m {
		//fmt.Printf("%v ---------------> %v\n", key, value)
		switch data := value.(type) {
		case string:
			str = data
			fmt.Printf("mao[%sa]的值类型为string,value = %s\n", key, str)

		}
	}
}
