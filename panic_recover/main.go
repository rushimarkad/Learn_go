//When a function encounters a panic, its execution is stopped,
//any deferred functions are executed and then the control returns to its caller.
//This process continues until all the functions of the current goroutine have returned at which point the program prints the panic message,
//followed by the stack trace and then terminates.

//=================================================

package main

import (
	"fmt"
)

//=====================================================

//Panic program

// func fullName(firstName *string, lastName *string) {
// 	if firstName == nil {
// 		panic("runtime error: first name cannot be nil")
// 	}
// 	if lastName == nil {
// 		panic("runtime error: last name cannot be nil")
// 	}
// 	fmt.Printf("%s %s\n", *firstName, *lastName)
// 	fmt.Println("returned normally from fullName")
// }

// func main() {
// 	firstName := "Elon"
// 	fullName(&firstName, nil)
// 	fmt.Println("returned normally from main")
// }

//==============================================

//It is possible to regain control of a panicking program using recover

//creates a panic due to out of bounds slice access.

// func slicePanic() {
// 	n := []int{5, 7, 4}
// 	fmt.Println(n[4]) //access index that is not present
// 	fmt.Println("normally returned from a")
// }
// func main() {
// 	slicePanic()
// 	fmt.Println("normally returned from main")
// }

//============================================

//defer

// func fullName(firstName *string, lastName *string) {
// 	defer fmt.Println("deferred call in fullName")
// 	if firstName == nil {
// 		panic("runtime error: first name cannot be nil")
// 	}
// 	if lastName == nil {
// 		panic("runtime error: last name cannot be nil")
// 	}
// 	fmt.Printf("%s %s\n", *firstName, *lastName)
// 	fmt.Println("returned normally from fullName")
// }

// func main() {
// 	defer fmt.Println("deferred call in main")
// 	firstName := "Elon"
// 	fullName(&firstName, nil)
// 	fmt.Println("returned normally from main")
// }

//=================================================

//Recovering from a Panic

//recover is a builtin function that is used to regain control of a panicking program.
//Recover is useful only when called inside deferred functions.

// func recoverFullName() {
// 	if r := recover(); r != nil {
// 		fmt.Println("recovered from ", r)
// 	}
// }

// func fullName(firstName *string, lastName *string) {
// 	defer recoverFullName()
// 	if firstName == nil {
// 		panic("runtime error: first name cannot be nil")
// 	}
// 	if lastName == nil {
// 		panic("runtime error: last name cannot be nil")
// 	}
// 	fmt.Printf("%s %s\n", *firstName, *lastName)
// 	fmt.Println("returned normally from fullName")
// }

// func main() {
// 	defer fmt.Println("deferred call in main")
// 	firstName := "Elon"
// 	fullName(&firstName, nil)
// 	fmt.Println("returned normally from main")
// }

//================================================

//one more example

// func recoverInvalidAccess() {
// 	if r := recover(); r != nil {
// 		fmt.Println("Recovered", r)
// 	}
// }

// func invalidSliceAccess() {
// 	defer recoverInvalidAccess()
// 	n := []int{5, 7, 4}
// 	fmt.Println(n[4])
// 	fmt.Println("normally returned from a")
// }

// func main() {
// 	invalidSliceAccess()
// 	fmt.Println("normally returned from main")
// }

//====================================================

//Getting Stack Trace after Recover

//If we recover from a panic, we lose the stack trace about the panic.
//There is a way to print the stack trace using the PrintStack function of the Debug package

// func recoverFullName() {
// 	if r := recover(); r != nil {
// 		fmt.Println("recovered from ", r)
// 		debug.PrintStack()
// 	}
// }

// func fullName(firstName *string, lastName *string) {
// 	defer recoverFullName()
// 	if firstName == nil {
// 		panic("runtime error: first name cannot be nil")
// 	}
// 	if lastName == nil {
// 		panic("runtime error: last name cannot be nil")
// 	}
// 	fmt.Printf("%s %s\n", *firstName, *lastName)
// 	fmt.Println("returned normally from fullName")
// }

// func main() {
// 	defer fmt.Println("deferred call in main")
// 	firstName := "Elon"
// 	fullName(&firstName, nil)
// 	fmt.Println("returned normally from main")
// }

//=============================================

//Panic, Recover and Goroutines
//Recover works only when it is called from the same goroutine which is panicking.
//It's not possible to recover from a panic that has happened in a different goroutine.

func recovery() {
	if r := recover(); r != nil {
		fmt.Println("recovered:", r)
	}
}

func sum(a int, b int) {
	defer recovery()
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	done := make(chan bool)
	go divide(a, b, done)
	//divide(a, b, done)	//using this statement will allow recovery
	<-done
}

func divide(a int, b int, done chan bool) {
	fmt.Printf("%d / %d = %d", a, b, a/b) //panics as cannot divide by zero
	done <- true

}

func main() {
	sum(5, 0)
	fmt.Println("normally returned from main")
}

//Here recovery is not done as recovery is done in sum but program panicks in divide go routine
