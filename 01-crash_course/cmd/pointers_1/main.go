package main

import (
	"fmt"
)

func main() {
	// Pointer variable
	// var p = 0x140000b2020
	// Normal variables
	// var i = 123
	// var p *int32 = new(int32)
	// var i int32
	//
	// fmt.Printf("The value p points to is: %v\n", *p)
	// fmt.Printf("The value if is: %v\n", i)

	// Resing a value
	// *p = 10
	// p = &i
	// *p = 444
	// fmt.Printf("The new value p points to is: %v\n", *p)
	// fmt.Printf("The value if is: %v\n", i)

	// Pointer in slices

	// var slice = []int32{1, 2, 3}
	// var sliceCopy = slice

	// sliceCopy[0] = 100

	// fmt.Println(slice)
	// fmt.Println(sliceCopy)

	// Without pointers

	var thing1 = [5]float64{1, 2, 3, 4, 5}
	fmt.Printf("\n\nThe memory location of the thing1 array is: %p", &thing1)
	var result [5]float64 = square(&thing1)
	fmt.Printf("\nThe result is: %v\n", result)
	fmt.Printf("The value of thing1 is: %v\n", thing1)

	// With pointers
	// The thing1 will be the same value with result
}

func square(thing2 *[5]float64) [5]float64 {
	fmt.Printf("\nThe memory location of the thing2 array is: %p\n", &thing2)
	for i := range thing2 {
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}
