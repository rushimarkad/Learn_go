// first class functions allows functions to be assigned to variables, passed as arguments to other functions
//and returned from other functions

package main

import (
	"fmt"
)

//==========================================

//Anonymous function assigned to variable

// func main() {
// 	a := func() { //assigned a function to the variable a
// 		fmt.Println("hello world first class function")
// 	}
// 	a()
// 	fmt.Printf("%T", a) //prints func()
// }

//==============================================

//Without assigned to variable

// func main() {
// 	func() {
// 		fmt.Println("hello world first class function")
// 	}() //Anonymous function call
// }

//================================================

//Passing args to anonymous func

// func main() {
// 	func(n string) {
// 		fmt.Println("Welcome", n)
// 	}("Gophers")
// }

//==============================================

//User defined function types

//Just like we define our own struct types, it is possible to define our own function types.

// type add func(a int, b int) int

// func main() {
// 	var a add = func(a int, b int) int { //define a variable a of type add and assign to it a function whose signature matches the type add
// 		return a + b
// 	}
// 	s := a(5, 6)
// 	fmt.Println("Sum", s)
// }

//================================================================

//Higher order functions
//function which does at least one of the following
//takes one or more functions as arguments
//returns a function as its result

//Passing functions as arguments to other functions

// func simple(a func(a, b int) int) {
// 	fmt.Println(a(60, 7))
// }

// func main() {
// 	f := func(a, b int) int {
// 		return a + b
// 	}
// 	simple(f)
// }

//===============================

//Returning functions from other functions

// func simple() func(a, b int) int {
// 	f := func(a, b int) int {
// 		return a + b
// 	}
// 	return f
// }

// func main() {
// 	s := simple()
// 	fmt.Println(s(60, 7))
// }

//===============================================

//Closures
// Closures are special case of anonymous functions that access the variables defined outside the body of the function.

// func main() {
// 	a := 5
// 	func() {
// 		fmt.Println("a =", a)
// 	}()
// }

//=================================================

//Every closure is bound to its own surrounding variable

// func appendStr() func(string) string {
// 	t := "Hello"
// 	c := func(b string) string {
// 		t = t + " " + b
// 		return t
// 	}
// 	return c
// }

// func main() {
// 	a := appendStr()
// 	b := appendStr()
// 	fmt.Println(a("World"))
// 	fmt.Println(b("Everyone"))

// 	fmt.Println(a("Gopher"))
// 	fmt.Println(b("!"))
// }

//======================================

//Practical use of first class functions
//create a program that filters a slice of students based on some criteria.

// type student struct {
// 	firstName string
// 	lastName  string
// 	grade     string
// 	country   string
// }

// func filter(s []student, f func(student) bool) []student {
// 	var r []student
// 	for _, v := range s {
// 		if f(v) == true {
// 			r = append(r, v)
// 		}
// 	}
// 	return r
// }

// func main() {
// 	s1 := student{
// 		firstName: "Naveen",
// 		lastName:  "Ramanathan",
// 		grade:     "A",
// 		country:   "India",
// 	}
// 	s2 := student{
// 		firstName: "Samuel",
// 		lastName:  "Johnson",
// 		grade:     "B",
// 		country:   "USA",
// 	}
// 	s := []student{s1, s2}
// 	f := filter(s, func(s student) bool {
// 		if s.grade == "B" { //to filter students with grade B
// 			return true
// 		}
// 		return false
// 	})
// 	fmt.Println(f)

// 	c := filter(s, func(s student) bool {
// 		if s.country == "India" { //To filter students having country India
// 			return true
// 		}
// 		return false
// 	})
// 	fmt.Println(c)
// }

//=======================================================

//multiply all integers in a slice by 5 and return the output

func iMap(s []int, f func(int) int) []int {
	var r []int
	for _, v := range s {
		r = append(r, f(v))
	}
	return r
}
func main() {
	a := []int{5, 6, 7, 8, 9}
	r := iMap(a, func(n int) int {
		return n * 5
	})
	fmt.Println(r)
}
