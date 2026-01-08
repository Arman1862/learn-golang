package main

import "fmt"

var globalInt int = 10

// group variable
var (
	globalName string = "Arman Rafardhan"
	globalAge  int16  = 17
)

func main() {
	fmt.Println(globalInt)
	something()

	// Multiple variable declaration
	// var year, name = 2026, "Fadil"
	// fmt.Println(year, name)
	
	// shortHand
	year := 2026
	name := "fadil"
	fmt.Println(year, name)

	// name := "julio" invalid because it's already declared
	name, age := "julio", 20
	fmt.Println(name, age)
}

func something() {
	fmt.Println("My name: " + globalName)
}