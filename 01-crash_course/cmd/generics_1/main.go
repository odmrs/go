package main

import (
	"fmt"
)

type sumSliceTypes interface {
	int | float32 | float64
}

func main() {
	// Generics are use to avoid repetation of creation of functions
	// Like create 5 functions just to change the type that function receive or return

	var intSlice = []int{1, 2, 3}
	fmt.Println(sumSlice[int](intSlice)) // don't need to specify the type to the sumSlice, the function can detect this easly

	var float32Slice = []float32{1, 2, 3}
	fmt.Println(sumSlice[float32](float32Slice)) // don't need to specify the type to the sumSlice, the function can detect this easly

	var float64Slice = []float64{1, 2, 3}
	fmt.Println(sumSlice[float64](float64Slice)) // don't need to specify the type to the sumSlice, the function can detect this easly
}

func sumSlice[T sumSliceTypes](slice []T) T {
	// This 'T' told to compiler that function can receive and return one int, float32 or float64
	var sum T
	for _, v := range slice {
		sum += v
	}
	return sum
}
