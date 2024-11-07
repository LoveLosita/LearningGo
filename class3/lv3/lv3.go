package main

import (
	"fmt"
	"math/rand"
	"time"
)

type ClientID struct {
	ConsumerID int
	Amount     int
}

func Timer(ToStock chan int) {
	for {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Second) //Restock goods in a random frequency
		ToStock <- 1
	}
}

func GenerateGoods(Buy chan ClientID, Stock chan int, ToStock chan int, Exit chan int) {
	Counts := 0
	var Amount ClientID
	time.Sleep(time.Duration(rand.Intn(20)))
	SameStock := rand.Intn(10) //Initial restock
	Stock <- SameStock
	for Counts <= 20 {
		select { //Listening for restock orders
		case <-ToStock:
			SameStock += rand.Intn(10) //each time restocking 0-10 goods
			Stock <- SameStock
			fmt.Printf("Restocked %d Goods!\n", SameStock)
			Counts += 1
		//Listening for purchases
		case Amount = <-Buy:
			if Amount.Amount <= SameStock {
				fmt.Printf("Consumer%d purchased %d goods!\n", Amount.ConsumerID, Amount.Amount)
				SameStock -= Amount.Amount
				Stock <- SameStock //update Stock
			} else {
				fmt.Printf("Sorry Consumer%d, the stock currently is %d, unable to meet your demend!\n", Amount.ConsumerID, SameStock)
			}
		default:
		}
	}
	close(Stock)
	Exit <- -1 //Exit signal
}

func Consumer1(Buy chan ClientID, Stock chan int) {
	for {
		var Purchase ClientID
		select {
		case Goods, ok := <-Stock:
			if !ok {
				return
			}
			time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
			if Goods == -1 {
				return
			} else {
				Purchase.ConsumerID = 1
				Purchase.Amount = rand.Intn(Goods)
				Buy <- Purchase
			}
		}
	}
}

func MeConsumer(Buy chan ClientID, Stock chan int) {
	for {
		var Purchase ClientID
		Purchase.ConsumerID = 2
		fmt.Scanf("%d", &Purchase.Amount)
		Buy <- Purchase
	}
}

func main() {
	Buy := make(chan ClientID, 100)
	Stock := make(chan int, 1)
	ToStock := make(chan int)
	Exit := make(chan int)
	go GenerateGoods(Buy, Stock, ToStock, Exit)
	go Consumer1(Buy, Stock)
	go MeConsumer(Buy, Stock)
	go Timer(ToStock)
	<-Exit
}
