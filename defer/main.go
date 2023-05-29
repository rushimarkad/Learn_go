//used to execute a function call just before the surrounding function where the defer statement is present returns

package main

import (
	"fmt"
	"sync"
)

//=====================================================
//Defer function

// func finished() {
// 	fmt.Println("Finished finding largest")
// }

// func largest(nums []int) {
// 	defer finished()
// 	fmt.Println("Started finding largest")
// 	max := nums[0]
// 	for _, v := range nums {
// 		if v > max {
// 			max = v
// 		}
// 	}
// 	fmt.Println("Largest number in", nums, "is", max)
// }

// func main() {
// 	nums := []int{78, 109, 2, 563, 300}
// 	largest(nums)
// }

//============================================

//Defer methods

// type person struct {
// 	firstName string
// 	lastName  string
// }

// func (p person) fullName() {
// 	fmt.Printf("%s %s", p.firstName, p.lastName)
// }

// func main() {
// 	p := person{
// 		firstName: "John",
// 		lastName:  "Smith",
// 	}
// 	defer p.fullName()
// 	fmt.Printf("Welcome ")
// }

//================================================

//Arguments evaluation

//The arguments of a deferred function are evaluated when the defer statement is executed and not when the actual function call is done.

// func printA(a int) {
// 	fmt.Println("value of a in deferred function", a)
// }
// func main() {
// 	a := 5
// 	defer printA(a)
// 	a = 10
// 	fmt.Println("value of a before deferred function call", a)

// }

//=======================================================

//Stack of defers

//When a function has multiple defer calls, they are pushed on to a stack and executed in Last In First Out (LIFO) order.

// func main() {
// 	name := "Naveen"
// 	fmt.Printf("Original String: %s\n", string(name))
// 	fmt.Printf("Reversed String: ")
// 	for _, v := range name {
// 		defer fmt.Printf("%c", v) //Prints in reverse order
// 	}
// }

//======================================

//Program without defer

// type rect struct {
// 	length int
// 	width  int
// }

// func (r rect) area(wg *sync.WaitGroup) {
// 	if r.length < 0 {
// 		fmt.Printf("rect %v's length should be greater than zero\n", r)
// 		wg.Done()
// 		return
// 	}
// 	if r.width < 0 {
// 		fmt.Printf("rect %v's width should be greater than zero\n", r)
// 		wg.Done()
// 		return
// 	}
// 	area := r.length * r.width
// 	fmt.Printf("rect %v's area %d\n", r, area)
// 	wg.Done()
// }

// func main() {
// 	var wg sync.WaitGroup
// 	r1 := rect{-67, 89}
// 	r2 := rect{5, -67}
// 	r3 := rect{8, 9}
// 	rects := []rect{r1, r2, r3}
// 	for _, v := range rects {
// 		wg.Add(1)
// 		go v.area(&wg)
// 	}
// 	wg.Wait()
// 	fmt.Println("All go routines finished executing")
// }

//==============================================

//Practical use of defer for above program

type rect struct {
	length int
	width  int
}

func (r rect) area(wg *sync.WaitGroup) {
	defer wg.Done() //Single Done instead of 3 as above
	if r.length < 0 {
		fmt.Printf("rect %v's length should be greater than zero\n", r)
		return
	}
	if r.width < 0 {
		fmt.Printf("rect %v's width should be greater than zero\n", r)
		return
	}
	area := r.length * r.width
	fmt.Printf("rect %v's area %d\n", r, area)
}

func main() {
	var wg sync.WaitGroup
	r1 := rect{-67, 89}
	r2 := rect{5, -67}
	r3 := rect{8, 9}
	rects := []rect{r1, r2, r3}
	for _, v := range rects {
		wg.Add(1)
		go v.area(&wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}
