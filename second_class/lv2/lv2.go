package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score []int
}

type Classroom struct {
	ClassName string
	Students  []*Student
}

func AddStudent(c *Classroom, s *Student) {
	c.Students = append(c.Students, s)
	fmt.Printf("Name:%s\nAge:%d\nScore:%v\n",
		c.Students[len(c.Students)-1].Name,
		c.Students[len(c.Students)-1].Age,
		c.Students[len(c.Students)-1].Score)
}

func UpdateScore(s *Student, score int) {
	s.Score = append(s.Score, score)
	fmt.Printf("scores:%v\n", s.Score)
}

func CalculateAverage(s *Student) float64 {
	sum := 0.0
	equal := 0.0
	for _, value := range s.Score {
		sum += float64(value)
	}
	equal = float64(sum) / float64(len(s.Score))
	return float64(equal)
}

func main() {
	stu1 := Student{Name: "1", Age: 17, Score: []int{100, 100, 101}}
	stu2 := Student{Name: "2", Age: 18, Score: []int{110, 105, 113}}
	class1 := Classroom{ClassName: "666", Students: []*Student{&stu1, &stu2}}
	newstu := Student{Name: "XiaoMing", Age: 18, Score: []int{100, 120, 130, 100, 100, 100}}
	AddStudent(&class1, &newstu)
	UpdateScore(class1.Students[0], 100)
	for _, value := range class1.Students {
		fmt.Printf("Name:%s,Score:%f\n", value.Name, CalculateAverage(value))
	}
}
