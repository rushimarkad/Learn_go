package main

import "fmt"

func main() {
	fmt.Print("Enter Your First Name: ")
	var first string
	fmt.Scanln(&first)
	fmt.Println("Your Name is: ", first)
}
