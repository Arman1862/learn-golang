package main

import "fmt"

func main() {
	// The default value of a string is ""
	var myString string
	fmt.Println("The value: " + myString)
	var myString2 string = "Hello"
	fmt.Println("The value: " + myString2)
	myString = "World"
	fmt.Println("The value: " + myString)

	// The default value of an int is 0
	var myInt int
	fmt.Println("The value: ", myInt)
	myInt = 20
	fmt.Println("The value: ", myInt)

	// The default value of a boolean is false
	var myBool bool
	fmt.Println("The value: ", myBool)
	myBool = true
	fmt.Println("The value: ", myBool)
}