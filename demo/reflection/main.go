package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// declaring a struct
type Person struct {
	// defining struct variables
	Name    string
	Age     int
	Country string
}

type Car struct {
	Company string
	Model   string
	Year    int
}

func main() {
	// defining a struct instance
	p := Person{
		Name:    "Rushi Markad",
		Age:     22,
		Country: "India",
	}

	c := Car{
		Company: "Toyota",
		Model:   "Innova",
		Year:    2018,
	}

	encode(p)
	encode(c)
}

func encode(data interface{}) {

	v := reflect.ValueOf(data) // Gives Values of struct
	fmt.Println("Value of : ", v)

	t1 := reflect.TypeOf(data) //Type of data
	fmt.Println("TypeOf: ", t1)

	t := v.Type() //Type of data
	fmt.Println("Type : ", t)

	k := t.Kind() //Actual type of interface
	fmt.Println("Kind : ", k)

	n := v.NumField() // NumField() method returns the number of fields in a struct.
	fmt.Println("Number of fields : ", n)

	// Create a map to hold the key-value pairs
	m := make(map[string]interface{})

	// Iterate over the fields of the struct
	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i) // Field(i int) method returns the reflect.Value of the ith field.
		value := v.Field(i).Interface()

		// Store the field name and value in the map
		m[field.Name] = value
	}

	fmt.Println(m)

	// Convert the map to JSON
	jsonData, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	fmt.Println(string(jsonData)) //is in a byte array, need to convert in string
	fmt.Println("")
}
