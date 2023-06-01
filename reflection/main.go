//Reflection is a ability of a program to inspect its variables and values at run time and find their type.

package main

import (
	"fmt"
	"reflect"
)

//===========================================================

//TO know datatype of variable

// func main() {
// 	i := 10
// 	fmt.Printf("%d %T", i, i)
// }

//========================================================

//General query creator

// type order struct {
// 	ordId      int
// 	customerId int
// }

// func createQuery(o order) string {
// 	i := fmt.Sprintf("insert into order values(%d, %d)", o.ordId, o.customerId)
// 	return i
// }

// func main() {
// 	o := order{
// 		ordId:      1234,
// 		customerId: 567,
// 	}
// 	fmt.Println(createQuery(o))
// }

//================================================

//The reflect package helps to identify the underlying concrete type and the value of a interface{} variable.
//There are two functions reflect.TypeOf() and reflect.ValueOf() which return the reflect.Type and reflect.Value respectively
//reflect.TypeOf(q)
//reflect.ValueOf(q)

// type order struct {
// 	ordId      int
// 	customerId int
// }

// func createQuery(q interface{}) {
// 	t := reflect.TypeOf(q)  //gets concrete type of the interface
// 	v := reflect.ValueOf(q) //gets value of interface
// 	k := t.Kind()           //gets the actual kind of the type
// 	fmt.Println("Type ", t)
// 	fmt.Println("Value ", v)
// 	fmt.Println("Kind ", k)

// }
// func main() {
// 	o := order{
// 		ordId:      456,
// 		customerId: 56,
// 	}
// 	createQuery(o)

// }

//=========================================

//NumField() and Field() methods

//NumField() method returns the number of fields in a struct.
//Field(i int) method returns the reflect.Value of the ith field.

// type order struct {
// 	ordId      int
// 	customerId int
// }

// func createQuery(q interface{}) {
// 	if reflect.ValueOf(q).Kind() == reflect.Struct { //check if it is a struct
// 		v := reflect.ValueOf(q)
// 		fmt.Println("Number of fields", v.NumField()) //gives num of fields
// 		for i := 0; i < v.NumField(); i++ {
// 			fmt.Printf("Field:%d type:%T value:%v\n", i, v.Field(i), v.Field(i))
// 		}
// 	}

// }
// func main() {
// 	o := order{
// 		ordId:      456,
// 		customerId: 56,
// 	}
// 	createQuery(o)
// }

//==============================================

//Int() and String() methods
//The methods Int and String help extract the reflect.Value as an int64 and string respectively.

// func main() {
// 	a := 56
// 	x := reflect.ValueOf(a).Int()
// 	fmt.Printf("type:%T value:%v\n", x, x)
// 	b := "Naveen"
// 	y := reflect.ValueOf(b).String()
// 	fmt.Printf("type:%T value:%v\n", y, y)

// }

//====================================================

//generalize our query creator and make it work on any struct

type order struct {
	ordId      int
	customerId int
}

type employee struct {
	name    string
	id      int
	address string
	salary  int
	country string
}

func createQuery(q interface{}) {
	if reflect.ValueOf(q).Kind() == reflect.Struct { //check if its struct
		t := reflect.TypeOf(q).Name() //gives name of struct
		query := fmt.Sprintf("insert into %s values(", t)
		v := reflect.ValueOf(q)
		for i := 0; i < v.NumField(); i++ {
			switch v.Field(i).Kind() {
			case reflect.Int: //extract the value of that field as int64 using the Int() method
				if i == 0 {
					query = fmt.Sprintf("%s%d", query, v.Field(i).Int())
				} else { //handle edge cases.
					query = fmt.Sprintf("%s, %d", query, v.Field(i).Int())
				}
			case reflect.String:
				if i == 0 {
					query = fmt.Sprintf("%s\"%s\"", query, v.Field(i).String())
				} else {
					query = fmt.Sprintf("%s, \"%s\"", query, v.Field(i).String())
				}
			default: //prevent the program from crashing when unsupported types are passed to the createQuery function
				fmt.Println("Unsupported type")
				return
			}
		}
		query = fmt.Sprintf("%s)", query)
		fmt.Println(query)
		return

	}
	fmt.Println("unsupported type")
}

func main() {
	o := order{
		ordId:      456,
		customerId: 56,
	}
	createQuery(o)

	e := employee{
		name:    "Naveen",
		id:      565,
		address: "Coimbatore",
		salary:  90000,
		country: "India",
	}
	createQuery(e)
	i := 90
	createQuery(i)

}
