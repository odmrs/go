package main // Told to compiler to start here

import (
	"fmt"
)

func main() {
	var intNum int16 = 32767
	intNum = intNum + 1 // Will overflow and return -32767
	fmt.Println(intNum)

	var floatNum float32 = 64.4
	fmt.Println(floatNum)

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	// Using strings
	fmt.Println("Same line with double quotes")
	fmt.Println(`Same text, diferent line
	like that`)
	fmt.Println("Concatenate" + " strings " + "using +")

	//GAMMA symbol
	// Rune is a type in golang that represente ONE character
	// Number of bytes in a text -> use len()
	// Number of character in a text use utf8.RuneInString()

	// Booleans

	fmt.Println(true)
	fmt.Println(false)

	// Default values

	// Almost number and rune is 0
	// String -> ''
	// Boolean -> false

	myVar := "text"
	fmt.Println(myVar)

	var var1, var2 int = 1, 2 // Multiple variables
	// or var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value" // Just can initialize with the values
}
