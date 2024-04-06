package main

import "fmt"

func main() {
	printMe("Something")
}

func printMe(printValue string) {
	fmt.Printf("Hello World, and a value: %v", printValue)
}
