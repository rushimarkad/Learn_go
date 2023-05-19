package main

import (
	"fmt"
)

//===========================

// func printBytes(s string) {
// 	fmt.Printf("Bytes: ")
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%x ", s[i])
// 	}
// }

// func printChars(s string) {
// 	fmt.Printf("Characters: ")
// 	// can also use runes
// 	for i := 0; i < len(s); i++ {
// 		fmt.Printf("%c ", s[i])
// 	}
// }

// func main() {
// 	name := "Hello World"
// 	fmt.Printf("String: %s\n", name)
// 	printChars(name)
// 	fmt.Printf("\n")
// 	printBytes(name)

// }

//=================================

// use utf8.RuneCountInString(s string) (n int) to find number of runes in a string

// func main() {
// 	word1 := "Señor"
// 	fmt.Printf("String: %s\n", word1)
// 	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
// 	fmt.Printf("Number of bytes: %d\n", len(word1))

// 	fmt.Printf("\n")
// 	word2 := "Pets"
// 	fmt.Printf("String: %s\n", word2)
// 	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
// 	fmt.Printf("Number of bytes: %d\n", len(word2))
// }

//==============================

//String comparison

// func compareStrings(str1 string, str2 string) {
// 	if str1 == str2 {	//compares two strings with == operator
// 		fmt.Printf("%s and %s are equal\n", str1, str2)
// 		return
// 	}
// 	fmt.Printf("%s and %s are not equal\n", str1, str2)
// }

// func main() {
// 	string1 := "Go"
// 	string2 := "Go"
// 	compareStrings(string1, string2)

// 	string3 := "hello"
// 	string4 := "world"
// 	compareStrings(string3, string4)

// }

//====================================================

//String concatenation

// func main() {
// 	string1 := "Go"
// 	string2 := "is awesome"
// 	result := string1 + " " + string2	//with + operator
// 	fmt.Println(result)
// }

//=====================================================

//using sprintf which formats a string according to the input format specifier

func main() {
	string1 := "Go"
	string2 := "is awesome"
	result := fmt.Sprintf("%s %s", string1, string2)
	fmt.Println(result)
}

//Strings are immutable
//To change it convert string to slice of runes and change and convrt again to string
