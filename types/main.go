package main

import (
	"fmt"
	"unsafe"
)

func main() {
	a := 89
	b := 95.67
	fmt.Println("value of a is", a, "and b is", b)
	fmt.Printf("type of a is %T, size of a is %d", a, unsafe.Sizeof(a))   //type and size of a
	fmt.Printf("\ntype of b is %T, size of b is %d", b, unsafe.Sizeof(b)) //type and size of b

	//complex function
	c1 := complex(5, 7)
	c2 := 8 + 27i
	cadd := c1 + c2
	fmt.Println("sum:", cadd)
	cmul := c1 * c2
	fmt.Println("product:", cmul)

	first := "Naveen"
	last := "Ramanathan"
	name := first + " " + last
	fmt.Println("My name is", name)

	// i := 55      //int
	// j := 67.8    //float64
	// sum := i + j //int + float64 not allowed
	// fmt.Println(sum)

	//Type COnversion
	i := 55           //int
	j := 67.8         //float64
	sum := i + int(j) //j is converted to int
	fmt.Println(sum)

}
