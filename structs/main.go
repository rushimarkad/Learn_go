package main

import (
	"fmt"
	"structs/computer"
)

func main() {
	spec := computer.Spec{
		Maker: "apple",
		Price: 50000,
		// model: "Kashmiri", 	// Cannot access the unexported field
	}
	fmt.Println("Maker:", spec.Maker)
	fmt.Println("Price:", spec.Price)
}
