package main

import (
	"fmt"
	"time"
)

func main() {
	var n int = 1000000                // Create a integer 64 because my machine is 64 bits
	var testSlice = []int{}            // Create a slice without lenght by default
	var testSlice2 = make([]int, 0, n) // And create a another slice with lenght 0 and capacity of n = 1000000

	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))
}

func timeLoop(slice []int, n int) time.Duration {
	var t0 = time.Now() // Create a object t0 with time of now

	for len(slice) > n {
		slice = append(slice, 1) // Add 1, to lenght of that slice
	}

	return time.Since(t0)
}
