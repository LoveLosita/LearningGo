package main

import (
	"fmt"
	"time"
)

var TotalTime time.Duration
var Duration time.Duration
var StartTime time.Time

type Flag struct {
	Num  int
	Time float64
}

func Timing(Option chan int, Leave chan int, Flags *[]Flag) {
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
			Pos := 0
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
					} else if CurrentOption == -1 { //-1 to exit
						Leave <- 1
						return
					} else {
						(*Flags)[Pos].Num = Pos + 1 //recording flags...
						(*Flags)[Pos].Time = Duration.Seconds()
						Pos += 1
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
	fmt.Printf(`Usage of the timer:
	1.Before starting:Press 1 to start.
	2.Started timing:Press 0 to reset, and press 1 to start again.
	Press any key except -1, 1, 2 to record grades(for instance I press 4).
	3.To pause:Press 2 to pause, and press 2 again to continue timing.
	4.When the flag recording ends, press -1 to exit, and the results will be shown automatically.
	`)
	Leave := make(chan int)
	Option := make(chan int)
	Flags := make([]Flag, 100)
	go Timing(Option, Leave, &Flags)
	go Inputing(Option)
	<-Leave
	for _, value := range Flags {
		if value.Num == 0 {
			break
		}
		fmt.Printf("No:%d,Time:%.3f\n", value.Num, value.Time)
	}
}
