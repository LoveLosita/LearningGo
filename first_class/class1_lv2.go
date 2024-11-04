package main

import (
	"fmt"
)

func main() {
	x := 0
	mod := ""
	y := 0
	flag := true
	option := "continue"
	for flag {
		fmt.Println("请输入第一个数字：")
		fmt.Scan(&x)
		fmt.Println("请输入运算符:")
		fmt.Scan(&mod)
		fmt.Println("请输入第二个数字：")
		fmt.Scan(&y)
		switch mod {
		case "*":
			fmt.Printf("%d*%d=%d", x, y, x*y)
		case "/":
			fmt.Printf("%d/%d=%f", x, y, float64(x)/float64(y))
		case "+":
			fmt.Printf("%d+%d=%d", x, y, x+y)
		case "-":
			fmt.Printf("%d-%d=%d", x, y, x-y)
		}
		fmt.Println("是否继续？exit退出")
		fmt.Scan(&option)
		if option == "exit" {
			flag = false
		}
	}
}
