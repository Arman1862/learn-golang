package main

import "fmt"

func main() {
	var smallFloat float32
	fmt.Println("Small float", smallFloat)
	smallFloat = 20.312234424
	fmt.Println("Small float", smallFloat)
	
	var bigFloat float64
	fmt.Println("Big float", bigFloat)
	bigFloat = 84.92473253837583573853857385
	fmt.Println("Big float", bigFloat)
	
	var myComplex complex128
	fmt.Println("Complex value", myComplex)
	myComplex = complex(bigFloat, bigFloat)
	fmt.Println("Complex value", myComplex)

	var myRealPart, myImaginaryPart float64
	myRealPart = real(myComplex)
	myImaginaryPart = imag(myComplex)
	fmt.Println("Real part", myRealPart)
	fmt.Println("Imaginary part", myImaginaryPart)
}