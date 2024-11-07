package main

import (
	"fmt"
)

func JiShu(Ji chan int, Ou chan int) {
	for i := 1; i < 100; i += 2 {
		<-Ji
		fmt.Println(i)
		Ou <- 1
	}
}

func OuShu(Ji chan int, Ou chan int, Finish chan int) {
	for j := 2; j <= 100; j += 2 {
		<-Ou
		fmt.Println(j)
		if j == 100 {
			Finish <- 1
		}
		Ji <- 1
	}
}

func main() {
	Ou := make(chan int)
	Ji := make(chan int)
	Finish := make(chan int)
	go JiShu(Ji, Ou)
	Ji <- 1
	go OuShu(Ji, Ou, Finish)
	<-Finish
}
