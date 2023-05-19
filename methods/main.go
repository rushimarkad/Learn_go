//method is just a function with a special receiver type between the func keyword and the method name
/*
Syntax:
	func (t Type) methodName(parameter list) {
	}
*/

package main

import (
	"fmt"
)

//==============================

// type Employee struct {
// 	name     string
// 	salary   int
// 	currency string
// }

// /*
// displaySalary() method has Employee as the receiver type
// */
// func (e Employee) displaySalary() {	//created a method displaySalary on Employee struct type
// 	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
// }

// func main() {
// 	emp1 := Employee{
// 		name:     "Sam Adolf",
// 		salary:   5000,
// 		currency: "$",
// 	}
// 	emp1.displaySalary() //Calling displaySalary() method of Employee type
// }

//===============================================

//Same method name with different receiver type

// type Rectangle struct {
// 	length int
// 	width  int
// }

// type Circle struct {
// 	radius float64
// }

// func (r Rectangle) Area() int {
// 	return r.length * r.width
// }

// func (c Circle) Area() float64 {
// 	return math.Pi * c.radius * c.radius
// }

// func main() {
// 	r := Rectangle{
// 		length: 10,
// 		width:  5,
// 	}
// 	fmt.Printf("Area of rectangle %d\n", r.Area())
// 	c := Circle{
// 		radius: 12,
// 	}
// 	fmt.Printf("Area of circle %f", c.Area())
// }

//=========================================================

//Reference type receiver

type Employee struct {
	name string
	age  int
}

/*
Method with value receiver
*/
func (e Employee) changeName(newName string) {
	e.name = newName
}

/*
Method with pointer receiver
*/
func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "Mark Andrew",
		age:  50,
	}
	fmt.Printf("Employee name before change: %s", e.name)
	e.changeName("Michael Andrew")
	fmt.Printf("\nEmployee name after change: %s", e.name)

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	e.changeAge(51)
	fmt.Printf("\nEmployee age after change: %d", e.age) //&e is replced by e
}
