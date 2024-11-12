package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//s := ""
	//fmt.Scanln(&s)
	//fmt.Println(s)
	// 打开标准输入
	r := bufio.NewReader(os.Stdin)
	// 读取标准输入的一行
	line, err := r.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(line)
	// 打开标准输出
	w := bufio.NewWriter(os.Stdout)
	// 写入一个字符串到标准输出
	_, err = w.WriteString("Hello, 世界\n")
	if err != nil {
		fmt.Println(err)
	}
	w.Flush()
}
