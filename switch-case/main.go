package main

import (
	"fmt"
)

//===================================

//Switch case

func main() {
	switch finger := 2; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: //default case
		fmt.Println("incorrect finger number")
	}
}

//=================================================
//Duplicate cases are not allowed
//=================================================

//multiple expression in case

// func main() {
// 	letter := "i"
// 	fmt.Printf("Letter %s is a ", letter)
// 	switch letter {
// 	case "a", "e", "i", "o", "u": //multiple expressions in case
// 		fmt.Println("vowel")
// 	default:
// 		fmt.Println("not a vowel")
// 	}
// }
