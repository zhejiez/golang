package nginxconf

import (
	"fmt"
	"os"
)

func Nginxconf(path, name string) {
	f, err := os.OpenFile(path, os.O_APPEND, 0644)
	var buf string
	if err != nil {
		fmt.Println("文件不存在: ", err)
		return
	} else {
		//写入文件
		buf = fmt.Sprintf("\nif  ($request_uri  ~* \"%s\") {\nreturn 403;\n}\n", name)
		//f.Seek(201,os.SEEK_SET)
		//f.WriteString(buf)
		// 从开头的偏移量开始写入内容
		n, _ := f.Seek(201, os.SEEK_SET)
		_, err = f.WriteAt([]byte(buf), n)
	}
	//使用完毕,关闭文件
	defer f.Close()
	//nginx_reload.Reload()
}
