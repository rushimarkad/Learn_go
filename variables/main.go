package main

import "fmt"

func main() {

	// var age int                   // variable declaration
	// fmt.Println("My age is", age) // gives 0
	// age = 29                      //assignment
	// fmt.Println("My age is", age)
	// age = 54 //reassignment
	// fmt.Println("My new age is", age)

	// var age int = 29 // variable declaration with initial value
	// fmt.Println("My age is", age)

	var age = 29 // type will be inferred
	fmt.Println("My age is", age)

	// var width, height int = 100, 50 //declaring multiple variables
	// fmt.Println("width is", width, "height is", height)

	var width, height = 100, 50 //"int" is dropped
	fmt.Println("width is", width, "height is", height)

	//Short hand Declaration
	count := 10
	fmt.Println("Count =", count)

	//constants
	const a = 55 //allowed
	//a = 89       //reassignment not allowed
	fmt.Println("a is : ", a)
}
