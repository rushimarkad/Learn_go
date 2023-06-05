package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {
	// Create a new JSON object
	jsonObj := gabs.New()

	// Set values for fields
	jsonObj.Set(42, "age")
	jsonObj.Set("John Doe", "name")
	jsonObj.Set(true, "isStudent")
	jsonObj.Set("example@example.com", "email")
	jsonObj.SetP("123 Main St", "address.street")
	jsonObj.SetP("New York", "address.city")

	// Get values from fields
	age := jsonObj.Path("age").Data().(int)
	name := jsonObj.Path("name").Data().(string)
	isStudent := jsonObj.Path("isStudent").Data().(bool)
	email := jsonObj.Path("email").Data().(string)
	street := jsonObj.Path("address.street").Data().(string)
	city := jsonObj.Path("address.city").Data().(string)

	// Print the retrieved values
	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Is Student:", isStudent)
	fmt.Println("Email:", email)
	fmt.Println("Street:", street)
	fmt.Println("City:", city)

	// Add a new field
	jsonObj.SetP("AFour", "company.name")

	// Check if a field exists
	if jsonObj.Exists("company.name") {
		fmt.Println("Field 'company.name' exists.")
	}

	// Remove a field
	jsonObj.Delete("address")

	// Stringify the JSON object
	jsonString := jsonObj.String()
	fmt.Println("JSON string:", jsonString)

	// Parse a JSON string
	parsedObj, err := gabs.ParseJSON([]byte(jsonString))
	if err != nil {
		fmt.Println("Failed to parse JSON:", err)
		return
	}

	// Get values from parsed object
	name = parsedObj.Path("name").Data().(string)
	company := jsonObj.Path("company.name").Data().(string)

	fmt.Println("Parsed Name:", name)
	fmt.Println("Parsed Company:", company)
}
