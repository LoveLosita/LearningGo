package main

import "fmt"

type Student struct {
	name  string
	age   int
	score int
}

func main() {
	stu1 := Student{name: "1", age: 18, score: 100}
	stu2 := Student{name: "2", age: 19, score: 112}
	stu3 := Student{name: "3", age: 20, score: 122}
	stus := []Student{stu1, stu2, stu3}
	for _, value := range stus {
		fmt.Printf("name:%s\n", value.name)
		fmt.Printf("age:%d\n", value.age)
		fmt.Printf("grade:%d\n", value.score)
	}
}
