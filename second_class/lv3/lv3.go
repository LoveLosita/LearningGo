package main

import (
	"fmt"
	"time"
)

type Task interface { //定义接口
	Execute() error
}

type PrintTask struct { //输出任务的结构体
	Message string
}

type CalculationTask struct { //计算任务结构体
	A int
	B int
}

type SleepTask struct { //休眠任务结构体，存储休眠时间
	Duration int
}

func (a PrintTask) Execute() error { //输出的方法，实现了执行接口
	fmt.Println(a.Message)
	return nil
}

func (b CalculationTask) Execute() error { //计算的方法，实现了执行接口
	C := b.B + b.A
	fmt.Println(C)
	return nil
}

func (c SleepTask) Execute() error { //休眠任务，实现了执行接口
	fmt.Println("开始休眠...")
	Duration := time.Duration(c.Duration)
	time.Sleep(Duration * time.Second)
	fmt.Println("休眠结束")
	return nil
}

type Scheduler struct { //任务结构体，存储一个个实现了该接口的任务
	Task []Task
}

func (d *Scheduler) AddTask(task Task) { //添加任务的方法
	d.Task = append(d.Task, task)
}

func (d *Scheduler) RunAll() { //运行全部的方法
	for _, value := range d.Task {
		err := value.Execute()
		if err != nil {
			fmt.Println("运行错误！")
		}
	}
}
func main() {
	Scheduler := Scheduler{}
	Scheduler.AddTask(PrintTask{Message: "学姐我不会啊！"})
	Scheduler.AddTask(CalculationTask{A: 1, B: 2})
	Scheduler.AddTask(SleepTask{Duration: 5})
	Scheduler.RunAll()
}
