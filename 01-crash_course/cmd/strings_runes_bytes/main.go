package main

import (
	"fmt"
)

func main() {
	var myString = []rune("résumé")
	// Will show:

	// r -> 114 = 01110010
	// é -> 233 = 11000011 + 10101001 from utf-32...
	// The len of this string will be 8, because the 'é' and the last 'é'
	var indexed = myString[0] // Will get the byte of first element
	fmt.Println(indexed)
	fmt.Printf("Value: %v, Type: %T\n", indexed, indexed) // Here the type is uint8
	fmt.Printf("The lenght of bytes from this string is: %v\n", len(myString))
	for i, v := range myString {
		fmt.Println(i, v)
	}

	var myRune = 'm'
	fmt.Printf("Value: %v, type: %T\n", string(myRune), myRune)
}

