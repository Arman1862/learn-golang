package main

import "fmt"

func main() {
	// uint8 is an unsigned 8-bit integer
	// unsigned is only positive value
	var smallPlusValue uint8
	smallPlusValue = 10
	fmt.Println(smallPlusValue)

	var smallValue int8
	smallValue = 22
	fmt.Println("small value =", smallValue)
	smallValue = -32
	fmt.Println("small value =", smallValue)

	// byte is alias for uint8
	// byte is mainly use to represent raw data
	var myByte byte
	myByte = 'A'
	fmt.Println("length :", myByte)
	myByte = 10
	fmt.Println("length :", myByte)

	// rune is alias for uint32 or int 32
	var myRune rune
	myRune = '🥸'
	fmt.Println("Length :", myRune)
}