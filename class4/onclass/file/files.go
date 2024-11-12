package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开一个文件
	f, err := os.Open("test.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close() // 记得关闭文件
	// 创建一个缓冲读取器
	r := bufio.NewReader(f)
	// 读取文件的一行
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	fmt.Println(line)
	// 读取文件的一个字节
	b, err := r.ReadByte()
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	// 创建一个文件
	f, err = os.Create("test2.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close() // 延时关闭
	// 创建一个缓冲写入器
	w := bufio.NewWriter(f)
	// 写入一个字符串到文件
	_, err = w.WriteString("Hello, 世界\n")
	if err != nil {
		panic(err)
	}
	// 刷新缓冲区， 将数据写入文件
	w.Flush()
	// 写入一个字节到文件
	err = w.WriteByte(65)
	if err != nil {
		panic(err)
	}
	w.Flush()
}
