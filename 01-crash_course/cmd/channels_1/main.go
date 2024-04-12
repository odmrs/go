package main

import (
	"fmt"
	"math/rand"
	"time"
)

// // Channels are like messagers between goroutines, avoid data race, hold and listen for data
//
//	func main() {
//		// var c = make(chan int) // Create a channel that will hold integers
//		c := make(chan int, 5) // Create a channel will hold 5 items
//		// // (channel <- 1) add one to the channel
//		// c <- 1
//		// var i = <-c
//		go process(c)
//		// fmt.Println(<-c) // This will pop out the value inside the channel
//		for i := range c {
//			fmt.Println(i)
//			time.Sleep(time.Second * 1)
//		}
//	}
//
//	func process(c chan int) {
//		defer close(c) // Will be sure of the channel was closed and the end
//		for i := 0; i < 5; i++ {
//			c <- i
//		}
//	}
var MAX_CHICKEN_PRICE float32 = 5
var MAX_TOFU_PRICE float32 = 3

func main() {
	t0 := time.Now()
	var chickenChannel = make(chan string)
	var tofuChannel = make(chan string)
	var websites = []string{"walmart.com", "costco.com", "sholefoods.com"}

	for i := range websites {
		go checkChickenPrices(websites[i], chickenChannel)
		go checkTofuPrices(websites[i], tofuChannel)
	}
	sendMessage(chickenChannel, tofuChannel)
	fmt.Printf("\nThe time to finish this process: %v\n", time.Since(t0))
}

func checkChickenPrices(website string, chickenChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var chickenPrice = rand.Float32() * 20
		if chickenPrice <= MAX_CHICKEN_PRICE {
			chickenChannel <- website
			break
		}
	}
}
func checkTofuPrices(website string, tofuChannel chan string) {
	for {
		time.Sleep(time.Second * 1)
		var tofuPrice = rand.Float32() * 20
		if tofuPrice <= MAX_TOFU_PRICE {
			tofuChannel <- website
			break
		}
	}
}

func sendMessage(chickenChannel chan string, tofuChannel chan string) {
	select {
	case website := <-chickenChannel:
		fmt.Printf("Found a deal on chicken at %s\n", website)
	case website := <-tofuChannel:
		fmt.Printf("Found a deal on tofu at %s\n", website)
	}
}
