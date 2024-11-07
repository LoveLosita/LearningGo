package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var mu sync.Mutex //As a lock
var Stock int     //Update Stock immediately

type ClientID struct {
	ConsumerID int
	Amount     int
}

func randomInt(min, max int) int {
	if min > max { //Check the equality of arguments
		fmt.Println("Invalid range: min should not be greater than max")
		return -1
	}
	if min == max {
		fmt.Println("Range is too small: min and max are equal")
		return min
	}
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

func GenerateGoods(ToStock chan int, Exit chan int, ToBuy chan int) {
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
			ToBuy <- Stock //Call the robot(Consumer1)to buy, and tell the amount
			mu.Unlock()
		default:
			time.Sleep(time.Millisecond * 20)
		}
	}
	Exit <- -1 //Exit signal
}

func Consumer1(Buy chan ClientID, ToBuy chan int) {
	for {
		Amount := <-ToBuy
		var Purchase ClientID
		time.Sleep(time.Duration(randomInt(1000, 3000)) * time.Microsecond) //Buy speed
		Purchase.ConsumerID = 1
		Purchase.Amount = randomInt(1, Amount)
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
	ToBuy := make(chan int)
	go GenerateGoods(ToStock, Exit, ToBuy)
	go ProcessPurchase(Buy)
	go Consumer1(Buy, ToBuy)
	go MeConsumer(Buy)
	go Timer(ToStock)
	<-Exit
}
