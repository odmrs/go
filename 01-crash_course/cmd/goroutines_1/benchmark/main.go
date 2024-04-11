package main

import (
	"fmt"
	"sync" // To use the wait group
	"time"
)
// Wait Group are counters
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}

func main() {
	t0 := time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)    // Set a start for WaitGroup
		go count() // To use goroutines, use before the function 'go' to set this as a goroutines
		// But in this way the program will finish instantanly, won't will wait the dbCall function finish first
	}
	wg.Wait() // Wait the counter of WaitGroup until are 0
	fmt.Printf("\nTotal executation time: %v\n", time.Since(t0))
}

func count(){
  var res int
  for i := 0; i < 100000; i++{
    res += 1
  fmt.Println(res)
  }
  wg.Done()
}
