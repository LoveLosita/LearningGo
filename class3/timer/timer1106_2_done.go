package main

import (
	"fmt"
	"time"
)

var TotalTime time.Duration
var Duration time.Duration
var StartTime time.Time

func Timing(Option chan int) {
	CurrentOption := <-Option
	FirstEntry := 1
	Exit := 0
	Pause := 0
	if CurrentOption == 0 || CurrentOption == 2 {
		fmt.Println("You haven't started timing yet!")
		return
	} else if CurrentOption != 1 {
		fmt.Println("Wrong Input!")
		return
	}
OuterLoop:
	for {
		if FirstEntry == 0 && Pause != 1 || Exit == 1 { //If not first time, waiting for command...
			CurrentOption = <-Option
		} else {
			FirstEntry = 0 //Lose its first time...
		}
		if Exit == 1 { //reset
			TotalTime = 0
			Exit = 0
			fmt.Println("Timer reseted!")
		}
		if CurrentOption == 1 { //start timing
			StartTime = time.Now()
			if Pause == 1 {
				Pause = 0
				StartTime = time.Now().Add(-TotalTime)
			}
			fmt.Println("Start timing!")
			for {
				select {
				case CurrentOption = <-Option: //Listening to commands while working...
					if CurrentOption == 0 {
						Exit = 1
						fmt.Println("Exiting!")
						goto OuterLoop
					} else if CurrentOption == 2 {
						Pause = 1
						goto OuterLoop
					} else {
						return
					}
				default:
				}
				Duration = time.Since(StartTime)
				fmt.Printf("%.2f\n", Duration.Seconds())
				time.Sleep(200 * time.Microsecond)
			}
		} else if Pause == 1 {
			fmt.Println("Paused!")
			fmt.Printf("%.2f Seconds passed!\n", Duration.Seconds())
			TotalTime = Duration
			CurrentOption = <-Option
			if CurrentOption == 0 { //reset
				Exit = 1
				continue
			} else if CurrentOption == 2 { //continue timing
				fmt.Println("Continue timing!")
				CurrentOption = 1
				continue
			} else {
				fmt.Println("Wrong Input!")
				return
			}
		}
	}
}

func Inputing(Option chan int) { //get input
	for {
		var Input int
		fmt.Scanf("%d", &Input)
		Option <- Input
	}
}

func main() {
	Option := make(chan int)
	go Timing(Option)
	go Inputing(Option)
	select {}
}
