package main

import "fmt"

// Give Each variable Data type
func calculateBill(price int, no int) int {
	var totalPrice = price * no
	return totalPrice
}

// Give Data type in common
// func calculateBill(price, no int) int {
// 	var totalPrice = price * no
// 	return totalPrice
// }

func main() {
	price, no := 90, 6
	totalPrice := calculateBill(price, no)
	fmt.Println("Total price is", totalPrice)
}

//===========================================

// func rectProps(length, width float64) (float64, float64) {
// 	var area = length * width
// 	var perimeter = (length + width) * 2
// 	return area, perimeter
// }

// func rectProps(length, width float64) (area, perimeter float64) {
// 	area = length * width
// 	perimeter = (length + width) * 2
// 	return //no explicit return value
// }

// func main() {
// 	area, perimeter := rectProps(10.8, 5.6)
// 	fmt.Printf("Area %f Perimeter %f", area, perimeter)
// }

//==============================================

// Blank Identifier

// func rectProps(length, width float64) (float64, float64) {
// 	var area = length * width
// 	var perimeter = (length + width) * 2
// 	return area, perimeter
// }
// func main() {
// 	area, _ := rectProps(10.8, 5.6) // perimeter is discarded
// 	fmt.Printf("Area %f ", area)
// }
