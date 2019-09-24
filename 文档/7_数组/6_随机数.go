package __数组

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	//设置种子,只需要一次
	//如果种子参数意义,每次运行的程序产生的随机数都一样
	rand.Seed(time.Now().UnixNano())	//以当前系统时间作为种子

	for i := 0; i < 5; i++ {
		//产生随机数
		//fmt.Println("rand = ", rand.Int())		//随机很大的数
		fmt.Println("rand = ", rand.Intn(100))
	}
}