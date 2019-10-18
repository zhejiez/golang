package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func main() {
	path := "/root/wh.super.com_vhost.conf"
	name := "asdzxc"
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("文件不存在: ", err)
		return
	} else {
		buf := fmt.Sprintf("\"2a if  (\\$request_uri  ~* \"%s\") {return 403;}\" %s", name, path)
		cmd := exec.Command("sed", "-i", buf)

		fmt.Println(cmd)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
			return
		} else if err := cmd.Start(); err != nil {
			fmt.Println("Error:The command is err,", err)
			return
		}
		//读取所有输出
		bytes, err := ioutil.ReadAll(stdout)
		if err != nil {
			fmt.Println("ReadAll Stdout:", err.Error())
			return
		}
		if err := cmd.Wait(); err != nil {
			fmt.Println("wait:", err.Error())
			return
		}
		fmt.Printf("stdout:\n\n %s", bytes)
	}
	defer f.Close()
}
