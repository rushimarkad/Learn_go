package main

import (
	"fmt"
)

//==================================

//Declaring the struct

// type Employee struct {
// 	firstName string
// 	lastName  string
// 	age       int
// 	salary    int
// }

// func main() {

// 	//creating struct specifying field names
// 	emp1 := Employee{
// 		firstName: "Sam",
// 		age:       25,
// 		salary:    500,
// 		lastName:  "Anderson",
// 	}

// 	//creating struct without specifying field names
// 	emp2 := Employee{"Thomas", "Paul", 29, 800}

// 	fmt.Println("Employee 1", emp1)
// 	fmt.Println("Employee 2", emp2)
// }

//======================

//Anonumous struct

// func main() {
// 	emp3 := struct {
// 		firstName string
// 		lastName  string
// 		age       int
// 		salary    int
// 	}{
// 		firstName: "Andreah",
// 		lastName:  "Nikola",
// 		age:       31,
// 		salary:    5000,
// 	}

// 	fmt.Println("Employee 3", emp3)
// }

//===============================

//Accessing individual fields of a struct

// type Employee struct {
// 	firstName string
// 	lastName  string
// 	age       int
// 	salary    int
// }

// func main() {
// 	emp6 := Employee{
// 		firstName: "Sam",
// 		lastName:  "Anderson",
// 		age:       55,
// 		salary:    6000,
// 	}
// 	fmt.Println("First Name:", emp6.firstName)
// 	fmt.Println("Last Name:", emp6.lastName)
// 	fmt.Println("Age:", emp6.age)
// 	fmt.Printf("Salary: $%d\n", emp6.salary)
// 	emp6.salary = 6500	//reassign value
// 	fmt.Printf("New Salary: $%d", emp6.salary)
// }

//When a struct is defined and it is not explicitly initialized with any value, the fields of the struct are assigned their zero values by default.

//===================================================================

//Pointer to struct

// type Employee struct {
// 	firstName string
// 	lastName  string
// 	age       int
// 	salary    int
// }

// func main() {
// 	emp8 := &Employee{
// 		firstName: "Sam",
// 		lastName:  "Anderson",
// 		age:       55,
// 		salary:    6000,
// 	}
// 	fmt.Println("First Name:", emp8.firstName)	//Go gives us the option to use emp8.firstName instead of the explicit dereference (*emp8).firstName
// 	fmt.Println("Age:", emp8.age)
// }

//============================================

//Anonumous fields : structs with fields that contain only a type without the field name

// type Person struct {
// 	string
// 	int
// }

// func main() {
// 	p1 := Person{
// 		string: "naveen",
// 		int:    50,
// 	}
// 	fmt.Println(p1.string)	//access with type of field
// 	fmt.Println(p1.int)
// }

//=========================================

//Nested Structs

// type Address struct {
// 	city  string
// 	state string
// }

// type Person struct {
// 	name    string
// 	age     int
// 	address Address
// }

// func main() {
// 	p := Person{
// 		name: "Naveen",
// 		age:  50,
// 		address: Address{
// 			city:  "Chicago",
// 			state: "Illinois",
// 		},
// 	}

// 	fmt.Println("Name:", p.name)
// 	fmt.Println("Age:", p.age)
// 	fmt.Println("City:", p.address.city)
// 	fmt.Println("State:", p.address.state)
// }

//=========================================

//Promoted fields

// type Address struct {
// 	city  string
// 	state string
// }
// type Person struct {
// 	name string
// 	age  int
// 	Address
// }

// func main() {
// 	p := Person{
// 		name: "Naveen",
// 		age:  50,
// 		Address: Address{
// 			city:  "Chicago",
// 			state: "Illinois",
// 		},
// 	}

// 	fmt.Println("Name:", p.name)
// 	fmt.Println("Age:", p.age)
// 	fmt.Println("City:", p.city)   //city is promoted field as field with type Address is anonymous therefore can access directly
// 	fmt.Println("State:", p.state) //state is promoted field
// }

//=============================================

//Comparison of struct

type name struct {
	firstName string
	lastName  string
}

func main() {
	name1 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name2 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	if name1 == name2 { //Structs are equal if values are equal
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	name3 := name{
		firstName: "Steve",
		lastName:  "Jobs",
	}
	name4 := name{
		firstName: "Steve",
	}

	if name3 == name4 {
		fmt.Println("name3 and name4 are equal")
	} else {
		fmt.Println("name3 and name4 are not equal")
	}
}
