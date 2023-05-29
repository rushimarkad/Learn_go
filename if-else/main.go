package main

import (
	"fmt"
)

// ===================================

//If else

func main() {
	num := 11
	if num%2 == 0 { //checks if number is even
		fmt.Println("the number", num, "is even")
	} else {
		fmt.Println("the number", num, "is odd")
	}
}

//=========================================

//If else if

// func main() {
// 	num := 99
// 	if num <= 50 {
// 		fmt.Println(num, "is less than or equal to 50")
// 	} else if num >= 51 && num <= 100 {
// 		fmt.Println(num, "is between 51 and 100")
// 	} else {
// 		fmt.Println(num, "is greater than 100")
// 	}
// }

//=========================================

// //with if assignment

// func main() {
// 	if num := 10; num%2 == 0 { //checks if number is even
// 		fmt.Println(num, "is even")
// 	} else {
// 		fmt.Println(num, "is odd")
// 	}
// }
