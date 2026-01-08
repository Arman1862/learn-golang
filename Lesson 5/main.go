package main

import "fmt"

func main() {
	var title string
	title = "basic"

	var fullTitle string
	fullTitle = "The title" + " " + title + ":"

	var myString string
	myString = "Hello World"
	fmt.Println(fullTitle + "\n" + myString)

	title = "/n escape"
	myString = "Hello\nWorld"
	fmt.Println(fullTitle + "\n" + myString)
	
	title = "raw format is backtick"
	myString = `Welcome
To
Golang`
	fmt.Println(fullTitle + "\n" + myString)

	var firstName, lastName, fullName string
	firstName = "Arman"
	lastName = "Rafardhan"

 	fullName = firstName + " " + lastName
	fmt.Println(fullName)
	fmt.Printf("%s %s\n", firstName, lastName)
	fullName = fmt.Sprintf("%s %s", firstName, lastName)
	fmt.Println(fullName)
}