package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex
var Stock int

type ClientID struct {
	ConsumerID int
	Amount     int
}

func randomInt(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func Timer(ToStock chan int) {
	for {
		time.Sleep(time.Duration(time.Duration(randomInt(5, 10)) * time.Second)) //Restock goods in a random frequency
		ToStock <- 1
	}
}

func ProcessPurchase(Buy chan ClientID) {
	for {
		select { //Listening purchases
		case Amount := <-Buy:
			mu.Lock()
			if Amount.Amount <= Stock {
				fmt.Printf("Consumer%d purchased %d goods!Current Stock:%d\n", Amount.ConsumerID, Amount.Amount, Stock-Amount.Amount)
				Stock -= Amount.Amount
			} else {
				fmt.Printf("Sorry Consumer%d, the stock currently is %d, unable to meet your demend!\n",
					Amount.ConsumerID, Stock)
			}
			mu.Unlock()
		default:
			time.Sleep(time.Millisecond * 20)
		}
	}
}

func GenerateGoods(ToStock chan int, Exit chan int) {
	Counts := 0
	mu.Lock()
	Stock = randomInt(5, 15) //initial Stock
	mu.Unlock()
	fmt.Printf("The initial stock is %d\n", Stock)
	for Counts <= 20 {
		select { //Listening for restock orders
		case <-ToStock: //Waiting for restock signal
			mu.Lock()
			Stock += randomInt(5, 15) //Each time restocking 5-15 goods
			fmt.Printf("Current Stock: %d Goods!\n", Stock)
			Counts += 1
			mu.Unlock()
		default:
			time.Sleep(time.Millisecond * 20)
		}
	}
	Exit <- -1 //Exit signal
}

func Consumer1(Buy chan ClientID) {
	for {
		var Purchase ClientID
		time.Sleep(time.Duration(randomInt(10, 20)) * time.Second)
		Purchase.ConsumerID = 1
		Purchase.Amount = randomInt(5, 10)
		Buy <- Purchase
	}
}

func MeConsumer(Buy chan ClientID) {
	for {
		var Purchase ClientID
		Purchase.ConsumerID = 2
		fmt.Scanf("%d", &Purchase.Amount)
		Buy <- Purchase
	}
}

func main() {
	Buy := make(chan ClientID, 100)
	ToStock := make(chan int)
	Exit := make(chan int)
	go GenerateGoods(ToStock, Exit)
	go ProcessPurchase(Buy)
	go Consumer1(Buy)
	go MeConsumer(Buy)
	go Timer(ToStock)
	<-Exit
}
