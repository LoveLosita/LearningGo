package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("欢迎使用本计时器！本计时器除了实现题目功能，还支持在暂停时结束哦！")
	option := 0
	flag := "NotPaused"
	for {
		var duration time.Duration
		fmt.Printf("请输入你的选择（0重置1开始2暂停或继续-1退出）")
		fmt.Scanf("%d", &option)
		if option == 1 {
			for {
				start_time := time.Now()
				fmt.Printf("请输入你的选择（0重置1开始结束2暂停或继续）")
				fmt.Scanf("%d", &option)
				if option == 0 {
					break
				} else if option == 1 {
					if flag == "Paused" {
						fmt.Printf("你本次计时了%f秒，计时结束\n", duration.Seconds())
						break
					}
					duration += time.Since(start_time)
					fmt.Printf("你本次计时了%f秒，计时结束\n", duration.Seconds())
					break
				} else if option == 2 {
					if flag == "NotPaused" {
						duration += time.Since(start_time)
						fmt.Printf("现在已经过去了%f秒，计时器已经暂停\n", duration.Seconds())
						flag = "Paused"
						continue
					} else {
						fmt.Printf("计时器继续计时\n")
						flag = "NotPaused"
					}
				}
			}
		} else if option == -1 {
			break
		} else {
			fmt.Println("你还没有开始计时器，请你先开始！")
		}
	}

}
