package main

import (
	"fmt"
	"time"
)

type Task interface {
	Execute() error
}

type PrintTask struct {
	Message string
}

type CalculationTask struct {
	A int
	B int
}

type SleepTask struct {
	Duration int
}

func (a PrintTask) Execute() error {
	fmt.Println(a.Message)
	return nil
}

func (b CalculationTask) Execute() error {
	C := b.B + b.A
	fmt.Println(C)
	return nil
}

func (c SleepTask) Execute() error {
	fmt.Println("开始休眠...")
	Duration := time.Duration(c.Duration)
	time.Sleep(Duration * time.Second)
	fmt.Println("休眠结束")
	return nil
}

type Scheduler struct {
	Task []Task
}

func (d *Scheduler) AddTask(task Task) {
	d.Task = append(d.Task, task)
}

func (d *Scheduler) RunAll() {
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
