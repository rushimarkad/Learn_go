package main

import (
	"fmt"
)

//==============================================

// array declaration

// func main() {
// 	var a [3]int //int array with length 3
// 	fmt.Println(a)	// prints 0 if not initialized
// }

//===============================================

// short hand declaration

// func main() {
// 	a := [3]int{12, 78, 50} // short hand declaration to create array
// 	fmt.Println(a)
// }

//============================================

//compiler decides the length

// func main() {
// 	a := [...]int{12, 78, 50} // ... makes the compiler determine the length
// 	fmt.Println(a)
// }

//========================================

//value remains unchanged when pass by value to functions

// func changeLocal(num [5]int) {
// 	num[0] = 55
// 	fmt.Println("inside function ", num)

// }
// func main() {
// 	num := [...]int{5, 6, 7, 8, 8}
//	fmt.Println("length of a is",len(num))	//To find the length of array
// 	fmt.Println("before passing to function ", num)
// 	changeLocal(num) //num is passed by value
// 	fmt.Println("after passing to function ", num)
// }

//======================================

//Looping with help of range

func main() {
	a := [...]float64{67.7, 89.8, 21, 78}
	sum := float64(0)
	for i, v := range a { //range returns both the index and value
		fmt.Printf("%d the element of a is %.2f\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}
