package main

import (
	"fmt"
)

//=======================================

// for loop

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d", i)
	}
}

// ======================================

//with break

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i > 5 {
// 			break //loop is terminated if i > 5
// 		}
// 		fmt.Printf("%d ", i)
// 	}
// 	fmt.Printf("\nline after for loop")
// }

//==========================================

//with continue

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		if i%2 == 0 {
// 			continue	//execution is skipped if i is 2
// 		}
// 		fmt.Printf("%d ", i)
// 	}
// }

// ========================================

// func main() {
// 	n := 5
// 	for i := 0; i < n; i++ {
// 		for j := 0; j <= i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

//======================================

// Infinite loop

// func main() {
// 	for {
// 		fmt.Println("Hello World")
// 	}
// }

//========================================

//while loop can be written as

// func main() {
// 	i := 0
// 	for i <= 10 { //semicolons are ommitted and only condition is present
// 		fmt.Printf("%d ", i)
// 		i += 2
// 	}
// }

//=========================================
