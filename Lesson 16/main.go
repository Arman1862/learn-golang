package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	a = [5]int{1,2,3,4,5}
	fmt.Println(a)

	b := [2][4]int{{1,2}, {1,2,3,4}}
	fmt.Println(len(b))
}