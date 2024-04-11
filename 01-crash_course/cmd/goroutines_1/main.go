package main

import (
	"fmt"
	"sync" // To use the wait group
	"time"
)

// var m = sync.Mutex{}
var m = sync.RWMutex{}

// Wait Group are counters
var wg = sync.WaitGroup{}

var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()

	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    // Set a start for WaitGroup
		go dbCall(i) // To use goroutines, use before the function 'go' to set this as a goroutines
		// But in this way the program will finish instantanly, won't will wait the dbCall function finish first
	}
	wg.Wait() // Wait the counter of WaitGroup until are 0
	fmt.Printf("\nTotal executation time: %v\n", time.Since(t0))
	fmt.Printf("\nThe result are: %v\n", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	// fmt.Println("The result from the database is: ", dbData[i])
	// m.Lock() // The goroutines will see that lock, wait until this mutex unlock and execute the append
	// This will avoid the cross write
	// result = append(result, dbData[i])
	// m.Unlock()
	// Mutext RWMutex
	save(dbData[i])
	log()
	wg.Done() // Will decrement the counter of waitgroup one
}

func save(result string) {
	m.Lock()
	results = append(results, result)
	m.Unlock()
}

func log() {
	// m.RLock()
	fmt.Printf("\nThe current results are: %v\n", results)
	// m.RUnlock()
}
