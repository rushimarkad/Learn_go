//Gabs is a small utility for dealing with dynamic or unknown JSON structures in Go.

package main

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
)

func main() {

	//Parsing and searching JSON
	jsonParsed, err := gabs.ParseJSON([]byte(`{
		"outer":{
			"inner":{
				"value1":10,
				"value2":22
			},
			"alsoInner":{
				"value1":20,
				"array1":[
					30, 40
				]
			}
		}
	}`))
	if err != nil {
		panic(err)
	}

	var value float64
	var ok bool

	value, ok = jsonParsed.Path("outer.inner.value1").Data().(float64)
	fmt.Println("Value :", value, "OK :", ok)
	// value == 10.0, ok == true

	value, ok = jsonParsed.Search("outer", "inner", "value1").Data().(float64)
	fmt.Println("Value :", value, "OK :", ok)
	// value == 10.0, ok == true

	value, ok = jsonParsed.Search("outer", "alsoInner", "array1", "1").Data().(float64)
	fmt.Println("Value :", value, "OK :", ok)
	// value == 40.0, ok == true

	gObj, err := jsonParsed.JSONPointer("/outer/alsoInner/array1/1")
	if err != nil {
		panic(err)
	}
	value, ok = gObj.Data().(float64)
	fmt.Println("Value :", value, "OK :", ok)
	// value == 40.0, ok == true

	value, ok = jsonParsed.Path("does.not.exist").Data().(float64)
	fmt.Println("Value :", value, "OK :", ok)
	// value == 0.0, ok == false

	exists := jsonParsed.Exists("outer", "inner", "value1")
	fmt.Println(exists)
	// exists == true

	exists = jsonParsed.ExistsP("does.not.exist")
	fmt.Println(exists)
	// exists == false

	//Iterating objects
	jsonParsed1, err := gabs.ParseJSON([]byte(`{"object":{"first":1,"second":2,"third":3}}`))
	if err != nil {
		panic(err)
	}

	// S is shorthand for Search
	for key, child := range jsonParsed1.S("object").ChildrenMap() {
		fmt.Printf("key: %v, value: %v\n", key, child.Data().(float64))
	}

	//Iterating arrays
	jsonParsed2, err := gabs.ParseJSON([]byte(`{"array":["first","second","third"]}`))
	if err != nil {
		panic(err)
	}

	for _, child := range jsonParsed2.S("array").Children() {
		fmt.Println(child.Data().(string))
	}

	//Children() will return all children of an array in order.
	//This also works on objects, however, the children will be returned in a random order.

	//Searching through arrays
	jsonParsed4, err := gabs.ParseJSON([]byte(`{"array":[{"value":1},{"value":2},{"value":3}]}`))
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonParsed4.Path("array.1.value").String())

	//Generating JSON
	jsonObj := gabs.New()
	// or gabs.Wrap(jsonObject) to work on an existing map[string]interface{}

	jsonObj.Set(10, "outer", "inner", "value")
	jsonObj.SetP(20, "outer.inner.value2")
	jsonObj.Set(30, "outer", "inner2", "value3")

	fmt.Println(jsonObj.String())
	//To pretty-print:
	fmt.Println(jsonObj.StringIndent("", "  "))

	//Generating Arrays
	jsonObj1 := gabs.New()

	jsonObj1.Array("foo", "array")
	// Or .ArrayP("foo.array")

	jsonObj1.ArrayAppend(10, "foo", "array")
	jsonObj1.ArrayAppend(20, "foo", "array")
	jsonObj1.ArrayAppend(30, "foo", "array")
	fmt.Println(jsonObj1.String())

	//Working with arrays by index:
	jsonObj2 := gabs.New()

	// Create an array with the length of 3
	jsonObj2.ArrayOfSize(3, "foo")

	jsonObj2.S("foo").SetIndex("test1", 0)
	jsonObj2.S("foo").SetIndex("test2", 1)

	// Create an embedded array with the length of 3
	jsonObj2.S("foo").ArrayOfSizeI(3, 2)

	jsonObj2.S("foo").Index(2).SetIndex(1, 0)
	jsonObj2.S("foo").Index(2).SetIndex(2, 1)
	jsonObj2.S("foo").Index(2).SetIndex(3, 2)
	fmt.Println(jsonObj2.String())

	//Converting back to JSON
	jsonParsedObj, _ := gabs.ParseJSON([]byte(`{
		"outer":{
			"values":{
				"first":10,
				"second":11
			}
		},
		"outer2":"hello world"
	}`))

	jsonOutput := jsonParsedObj.String()
	fmt.Println(jsonOutput)
	// Becomes `{"outer":{"values":{"first":10,"second":11}},"outer2":"hello world"}`

	//And to serialize a specific segment is as simple as:
	jsonOutput1 := jsonParsedObj.Search("outer").String()
	fmt.Println(jsonOutput1)
}
