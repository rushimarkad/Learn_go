package main

import (
	"fmt"
)

//=========================================

//Creating map

// func main() {
// 	employeeSalary := make(map[string]int)
// 	fmt.Println(employeeSalary)
// }

//======================================

//Adding items to a map

// func main() {
// 	employeeSalary := make(map[string]int)
// 	employeeSalary["steve"] = 12000
// 	employeeSalary["jamie"] = 15000
// 	employeeSalary["mike"] = 9000
// 	fmt.Println("employeeSalary map contents:", employeeSalary)
// }

func main() {
	// //Initialization during declaration
	// employeeSalary := map[string]int{
	// 	"steve": 12000,
	// 	"jamie": 15000,
	// }
	// employeeSalary["mike"] = 9000
	// fmt.Println("employeeSalary map contents:", employeeSalary)

	// //retriving value of key from map
	// employee := "jamie"
	// salary := employeeSalary[employee]
	// fmt.Println("Salary of", employee, "is", salary)

	// //To check if key exists
	// newEmp := "joe"
	// value, ok := employeeSalary[newEmp]
	// if ok == true {
	// 	fmt.Println("Salary of", newEmp, "is", value)
	// 	return
	// }
	// fmt.Println(newEmp, "not found")

	// //Iterate over map
	// fmt.Println("Contents of the map")
	// for key, value := range employeeSalary {
	// 	fmt.Printf("employeeSalary[%s] = %d\n", key, value)
	// }

	// //delete from map
	// fmt.Println("map before deletion", employeeSalary)
	// delete(employeeSalary, "steve")
	// fmt.Println("map after deletion", employeeSalary)

	// // Length of map

	// fmt.Println("length is", len(employeeSalary))

	//==============================================================
	// maps are of reference types
	//==============================================================

	// //map of structs
	// type employee struct {
	// 	salary  int
	// 	country string
	// }

	// emp1 := employee{
	// 	salary:  12000,
	// 	country: "USA",
	// }
	// emp2 := employee{
	// 	salary:  14000,
	// 	country: "Canada",
	// }
	// emp3 := employee{
	// 	salary:  13000,
	// 	country: "India",
	// }
	// employeeInfo := map[string]employee{
	// 	"Steve": emp1,
	// 	"Jamie": emp2,
	// 	"Mike":  emp3,
	// }

	// for name, info := range employeeInfo {
	// 	fmt.Printf("Employee: %s Salary:$%d  Country: %s\n", name, info.salary, info.country)
	// }

	//===================================================================

	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike":  9000,
	}
	fmt.Println("Original employee salary", employeeSalary)
	modified := employeeSalary
	modified["mike"] = 18000
	fmt.Println("Employee salary changed", employeeSalary)

}
