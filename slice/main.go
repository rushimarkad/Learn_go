package main

import (
	"fmt"
)

//=====================================

//Creating slice

// func main() {
// 	a := [5]int{76, 77, 78, 79, 80}
// 	var b []int = a[1:4] //creates a slice from a[1] to a[3]
// 	fmt.Println(b)
// }

//======================================

//another way

// func main() {
//     c := []int{6, 7, 8} //creates and array and returns a slice reference
//     fmt.Println(c)
// }

//============================================

//value by ref i.e change in slice will reflect in array

// func main() {
// 	numa := [3]int{78, 79, 80}
// 	nums1 := numa[:] //creates a slice which contains all elements of the array
// 	nums2 := numa[:]
// 	fmt.Println("array before change 1", numa)
// 	nums1[0] = 100
// 	fmt.Println("array after modification to slice nums1", numa)
// 	nums2[1] = 101
// 	fmt.Println("array after modification to slice nums2", numa)
// }

//==============================================

// Size and Capacity of slice

// func main() {
// 	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
// 	fruitslice := fruitarray[1:3]
// 	fmt.Printf("length of slice %d capacity %d \n", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and capacity is 6
// 	fruitslice = fruitslice[:cap(fruitslice)]                                         //re-slicing furitslice till its capacity
// 	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
// }

//zero value of slice is nill.

//=============================================

//Creating slice using make

// func main() {
// 	i := make([]int, 5, 5)
// 	fmt.Println(i)
// }

//======================================

//Appending to slice

// func main() {
// 	cars := []string{"Ferrari", "Honda", "Ford"}
// 	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
// 	cars = append(cars, "Toyota")
// 	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
// }

//======================================

//apeending one slice to another

func main() {
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
}
