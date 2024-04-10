package main

import (
	"fmt"
)

func main() {
	// ARRAYS
	// Arrays are fixed length, same type and indexable
	var intArr [3]int32 // fixed lenght -> 3 | Same type -> int32 Starting with [0, 0, 0]
	intArr[0] = 124
	intArr[1] = 321
	intArr[2] = 231

	// Or initialize this array like that
	var intArr2 [3]int32 = [3]int32{1, 2, 3}

	// Or with short asing operator

	intArr3 := [3]int32{1, 2, 3}
	// fmt.Println(intArr[0]) // Indexable
	// fmt.Println(intArr[1:3])

	// Using & to show the address of memory
	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	fmt.Println(intArr2)
	fmt.Println(intArr3)

	// SLICES

	// Using slices -> wrap arrays to give a more general, powerful, andconvenient interface to sequences of data
	var intSlice []int32 = []int32{4, 5, 6}
	// Cap -> show the capacity of slice have
	fmt.Printf("The lenght is %v, with capacity %v\n", len(intSlice), cap(intSlice))

	// Can append a value using append(nameOfVariable, value)
	intSlice = append(intSlice, 7)
	fmt.Printf("The lenght is %v, with capacity %v\n", len(intSlice), cap(intSlice))

	intSlice2 := []int32{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)

	// Or can create a slice using the make function
	var intSlice3 []int32 = make([]int32, 3, 8) // A slice of type int32 with the start lenght ':3', and the capacity '8'
	fmt.Println(intSlice3)

	// MAP Using map is a data structure to use "key": "value"
	// using make function

	var myMap map[string]uint8 = make(map[string]uint8)
	// That means that keys are strings, and values are uint8
	fmt.Println(myMap)

	// Or always initialize that map like that

	var myMap2 = map[string]uint8{"Adan": 23, "Marcos": 19}
	fmt.Println(myMap2)

	// To get the value:
	fmt.Println(myMap2["Marcos"]) // 19

	// If you try access a key that not exist will return the default value of this type uint8 -> 0

	// But the map return two values, the value and the boolean, if that item exist or not
	var age, ok = myMap2["Marcos"]
	if ok {
		fmt.Printf("Exist, the age is %v\n", age)
	} else {
		fmt.Println("Invalid name")
	}

	// can delete a value too, like that

	// delete(myMap2, "Adan")
	// fmt.Println(myMap2)

	// FOR

	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}
}
