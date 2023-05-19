package main

import (
	"fmt"
)

//======================

//Pointer declaration

// func main() {
// 	b := 255
// 	var a *int = &b
// 	fmt.Printf("Type of a is %T\n", a)
// 	fmt.Println("address of b is", a)
// }

//================================

//Zero value of pointer is nil

// func main() {
// 	a := 25
// 	var b *int
// 	if b == nil {
// 		fmt.Println("b is", b)
// 		b = &a
// 		fmt.Println("b after initialization is", b)
// 	}
// }

//=============================================

//new

// func main() {
// 	size := new(int)
// 	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
// 	*size = 85
// 	fmt.Println("New size value is", *size)
// }

//====================================

//Dereferece pointer

// func main() {
// 	b := 255
// 	a := &b
// 	fmt.Println("address of b is", a)
// 	fmt.Println("value of b is", *a)
// 	*a++
// 	fmt.Println("new value of b is", b)
// }

//===================================

//Pass pointer to function

// func change(val *int) {
// 	*val = 55
// }
// func main() {
// 	a := 58
// 	fmt.Println("value of a before function call is", a)
// 	b := &a
// 	change(b)
// 	fmt.Println("value of a after function call is", a)
// }

//==========================================

//Return pointer from function

func hello() *int {
	i := 5
	return &i
}
func main() {
	d := hello()
	fmt.Println("Value of d", *d)
}

//======================================
