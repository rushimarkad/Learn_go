// encoding/json package offers API functions for generating JSON documents from Go objects and vice versa.
//Also, it allows you to customize the JSON-to-Go and Go-to-JSON translation process.

//The Marshal() function in package encoding/json is used to encode the go object or data into JSON.
//func Marshal(v interface{}) ([]byte, error)

// The Unmarshal() function in package encoding/json is used to unpack or decode the go object data from JSON to struct.
// func Unmarshal(data []byte, v interface{}) error

package main

import (
	"encoding/json"
	"fmt"
)

//===============================================

// encode to JSON from a map

// func main() {
// 	fileCount := map[string]int{
// 		"cpp": 10,
// 		"js":  8,
// 		"go":  10,
// 	}
// 	bytes, _ := json.Marshal(fileCount) //is in a byte array, need to convert in string
// 	fmt.Println(string(bytes))
// }

//===========================================

// encode to JSON from a struct

// type Book struct {
// 	Title  string
// 	Author string
// 	Year   int
// }

// func main() {
// 	myBook := Book{
// 		Title:  "Hello Golang",
// 		Author: "John Mike",
// 		Year:   2021,
// 	}
// 	bytes, _ := json.Marshal(myBook)
// 	fmt.Println(string(bytes))
// }

//=============================================

// Marshaling complex objects

// If you try to encode integer arrays, string arrays, and primitive variables,  Go will produce simple JSON structures for those elements.

// type Seller struct {
// 	Id          int
// 	Name        string
// 	CountryCode string
// }
// type Product struct {
// 	Id     int
// 	Name   string
// 	Seller Seller
// 	Price  int
// }

// func main() {
// 	products := []Product{
// 		Product{
// 			Id:     50,
// 			Name:   "Writing Book",
// 			Seller: Seller{1, "ABC Company", "US"},
// 			Price:  100,
// 		},
// 		Product{
// 			Id:     51,
// 			Name:   "Kettle",
// 			Seller: Seller{20, "John Store", "DE"},
// 			Price:  500,
// 		},
// 	}
// 	bytes, _ := json.Marshal(products)
// 	fmt.Println(string(bytes))
// }

//=============================================

//The output JSON structure’s keys always start with an upper case English letter — how can we rename JSON fields?
//When we encode large complex structures, the output becomes literally unreadable — how can we prettify the JSON output?

// ==============================================

// Renaming fields

// encoding/json package allows developers to rename JSON fields as they wish via JSON struct tags.

// type Seller struct {
// 	Id          int    `json:"id"`
// 	Name        string `json:"name"`
// 	CountryCode string `json:"country_code"`
// }

// type Product struct {
// 	Id     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Seller Seller `json:"seller"`
// 	Price  int    `json:"price"`
// }

// func main() {
// 	book := Product{
// 		Id:     50,
// 		Name:   "Writing Book",
// 		Seller: Seller{1, "ABC Company", "US"},
// 		Price:  100,
// 	}
// 	bytes, _ := json.Marshal(book)
// 	fmt.Println(string(bytes))

// 	//Generating JSON with indentation (pretty-print)
// 	bytes1, _ := json.MarshalIndent(book, "", "\t")
// 	fmt.Println(string(bytes1))
// }

// above code uses struct tags to rename each exported field

//====================================================

// Ignoring specific fields from JSON output

//use json:”-” as the tag, the related struct field won’t be used for encoding.
//if we use ,omitempty inside the struct tag name string, the related field won’t be used for encoding if the value is empty.

//following code omits the product identifier for encoding. Also, it omits empty country code values from the output.

// type Seller struct {
// 	Id          int    `json:"id"`
// 	Name        string `json:"name"`
// 	CountryCode string `json:"country_code,omitempty"` // field won’t be used for encoding if the value is empty.
// }
// type Product struct {
// 	Id     int    `json:"-"` //field won’t be used for encoding
// 	Name   string `json:"name"`
// 	Seller Seller `json:"seller"`
// 	Price  int    `json:"price"`
// }

// func main() {
// 	products := []Product{
// 		Product{
// 			Id:     50,
// 			Name:   "Writing Book",
// 			Seller: Seller{Id: 1, Name: "ABC Company", CountryCode: "US"},
// 			Price:  100,
// 		},
// 		Product{
// 			Id:     51,
// 			Name:   "Kettle",
// 			Seller: Seller{Id: 20, Name: "John Store"},
// 			Price:  500,
// 		},
// 	}
// 	bytes, _ := json.MarshalIndent(products, "", "\t")
// 	fmt.Println(string(bytes))
// }

//===========================================================================

//Unmarshaling: Converting JSON to Go objects

//============================================================================

//Unmarshaling simple JSON structures

// Decode from json to struct

// type Window struct {
// 	Width  int    `json:"width"`
// 	Height int    `json:"height"`
// 	Title  string `json:"title"`
// }

// func main() {
// 	jsonInput := `{
//         "width": 500,
//         "height": 200,
//         "title": "Hello Go!"
//     }`
// 	var window Window
// 	err := json.Unmarshal([]byte(jsonInput), &window)	//jsonInput variable holds the JSON content as a multiline string. Therefore, we had to convert it to a bytes slice before passing it to the Unmarshal function with the byte[]()

// 	if err != nil {
// 		fmt.Println("JSON decode error!")
// 		return
// 	}

// 	fmt.Println(window) //prints the struct
// }

//=======================================================

// decode JSON structures to Go maps

// func main() {
// 	jsonInput := `{
//         "apples": 10,
//         "mangos": 20,
//         "grapes": 20
//     }`
// 	var fruitBasket map[string]int
// 	err := json.Unmarshal([]byte(jsonInput), &fruitBasket)

// 	if err != nil {
// 		fmt.Println("JSON decode error!")
// 		return
// 	}

// 	fmt.Println(fruitBasket) //OP= map[apples:10 grapes:20 mangos:20]
// }

//=====================================================

//Unmarshaling complex data structures

// type Product struct {
// 	Id     int    `json:"id"`
// 	Name   string `json:"name"`
// 	Seller struct {
// 		Id          int    `json:"id"`
// 		Name        string `json:"name"`
// 		CountryCode string `json:"country_code"`
// 	} `json:"seller"`
// 	Price int `json:"price"`
// }

// func main() {
// 	jsonInput := `[
//     {
//         "id":50,
//         "name":"Writing Book",
//         "seller":{
//             "id":1,
//             "name":"ABC Company",
//             "country_code":"US"
//         },
//         "price":100
//     },
//     {
//         "id":51,
//         "name":"Kettle",
//         "seller":{
//             "id":20,
//             "name":"John Store",
//             "country_code":"DE"
//         },
//         "price":500
//     }]
//     `
// 	var products []Product
// 	err := json.Unmarshal([]byte(jsonInput), &products)

// 	if err != nil {
// 		fmt.Println("JSON decode error!")
// 		return
// 	}

// 	fmt.Println(products)
// 	// OP= [{50 Writing Book {1 ABC Company US} 100} {51 Kettle {20 John Store DE} 500}]
// }

//============================================================

// Reading JSON files from the filesystem

// type Config struct {
// 	Timeout     float32
// 	PluginsPath string
// 	Window      struct {
// 		Width  int
// 		Height int
// 		X      int
// 		Y      int
// 	}
// }

// func main() {

// 	bytes, err := ioutil.ReadFile("config.json") //reads the JSON file content as bytes

// 	if err != nil {
// 		fmt.Println("Unable to load config file!")
// 		return
// 	}

// 	var config Config
// 	err = json.Unmarshal(bytes, &config) //decodes data records to the Config struct.

// 	if err != nil {
// 		fmt.Println("JSON decode error!")
// 		return
// 	}

// 	fmt.Println(config) // OP= {50.3 ~/plugins/ {500 200 500 500}}
// }

//==========================================================

//Writing JSON files to the filesystem

// type Window struct {
// 	Width  int `json:"width"`
// 	Height int `json:"height"`
// 	X      int `json:"x"`
// 	Y      int `json:"y"`
// }
// type Config struct {
// 	Timeout     float32 `json:"timeout"`
// 	PluginsPath string  `json:"pluginsPath"`
// 	Window      Window  `json:"window"`
// }

// func main() {
// 	config := Config{
// 		Timeout:     40.420,
// 		PluginsPath: "~/plugins/etc",
// 		Window:      Window{500, 200, 20, 20},
// 	}
// 	bytes, _ := json.MarshalIndent(config, "", "  ")
// 	ioutil.WriteFile("config1.json", bytes, 0644)
// }

//================================================================

//Custom marshaling

//you can customize the JSON encoding process via the custom marshaling feature in the Go json package
//Hide email addresses from below

// type Person struct {
// 	Name  string `json:"name"`
// 	Age   int    `json:"age"`
// 	Email string `json:"-"`
// }

// func main() {
// 	persons := []Person{
// 		Person{"James Henrick", 25, "james.h@gmail.com"},
// 		Person{"David Rick", 30, "rick.dvd@yahoo.com"},
// 	}
// 	bytes, _ := json.MarshalIndent(persons, "", "  ")
// 	fmt.Println(string(bytes))
// }
// func (p *Person) MarshalJSON() ([]byte, error) {
// 	type PersonAlias Person
// 	return json.Marshal(&struct {
// 		*PersonAlias
// 		Email string `json:"email"`
// 	}{
// 		PersonAlias: (*PersonAlias)(p),
// 		Email:       strings.Repeat("*", 4) + "@mail.com", // alter email
// 	})
// }

//========================================================

//Custom unmarshaling

//field says the temperature in Kelvin, but you need to store the specific value in Celsius.

type Config struct {
	FunctionName string
	Temperature  float32
}

func main() {
	jsonInput := `{
        "functionName": "triggerModule",
        "temperature": 4560.32
    }`
	var config Config
	err := json.Unmarshal([]byte(jsonInput), &config)

	if err != nil {
		fmt.Println("JSON decode error!")
		return
	}

	fmt.Println(config) // {triggerModule 4287.17}
}
func (c *Config) UnmarshalJSON(data []byte) error {
	type ConfigAlias Config
	tmp := struct {
		Temperature float32
		*ConfigAlias
	}{
		ConfigAlias: (*ConfigAlias)(c),
	}
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	c.Temperature = tmp.Temperature - 273.15
	return nil
}
