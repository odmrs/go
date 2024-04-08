package main

import (
	"errors"
	"fmt"
)

func main() {
	var result, remainder, err = intDivision(300, 0)
	if err != nil {
		fmt.Printf(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result the integer division is %v\n", result)
	} else {
		fmt.Printf("The result of the integer division is %v\nAnd the result of remainder was %v\n", result, remainder)
	}

	// Or using switch

	switch {
	case err != nil:
		fmt.Printf(err.Error())
	case remainder == 0:
		fmt.Printf("The result the integer division is %v\n", result)
	default:
		fmt.Printf("The resultof the integer division is %v and the result of remainder is %v\n", result, remainder)
	}
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error

	if denominator == 0 {
		err = errors.New("Can't divide a value by 0")
		return 0, 0, err
	}

	// func nameOfFunc(nameOfPlace int) typeOfReturn
	// Return multiple values:
	// func nameOfFunc(value int) (int, int){}
	var result = numerator / denominator
	var remainder = numerator % denominator
	return result, remainder, err
}
