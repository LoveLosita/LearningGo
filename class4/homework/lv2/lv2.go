package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type dateAndEvent struct {
	date  string
	event string
}

func main() {
	events := []dateAndEvent{}
	f, err := os.Open("events.txt")
	if err != nil {
		panic(err)
	}
	defer func(f *os.File) {
		_ = f.Close()
	}(f) // 记得关闭文件
	r := bufio.NewReader(f) // 创建一个缓冲读取器
	for {
		// ReadString 以换行符为分隔符读取一行
		line, err := r.ReadString('\n')
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Println("Error reading file:", err)
			}
			break
		}
		tempSlice := strings.Split(line, " ")
		tempType := dateAndEvent{date: tempSlice[0], event: tempSlice[1]}
		events = append(events, tempType)
	}
	minTimeGap := 365 * 24 * time.Hour
	minTimeEvent := ""
	for j := 1; j < len(events); j++ {
		t := events[j].date
		date, _ := time.Parse("2006-01-02", t)
		TimeGap := date.Sub(time.Now())
		if TimeGap < minTimeGap && TimeGap > 0 {
			minTimeGap = TimeGap
			minTimeEvent = events[j].event
		}
	}
	fmt.Printf("最近的一个事件是：%s\n还有%v天", minTimeEvent, minTimeGap.Hours()/24)
	return
}
