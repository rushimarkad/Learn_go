//solve race conditions using mutexes and channels.
//When a program runs concurrently, the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time.
//This section of code that modifies shared resources is called critical section.
//This is solved by mutex
//A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point in time
//to prevent race conditions from happening.

//Mutex is available in the sync package. There are two methods defined on Mutex namely Lock and Unlock.
//Any code that is present between a call to Lock and Unlock will be executed by only one Goroutine

package main

import (
	"fmt"
	"sync"
)

//=======================================

//program with race condition

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}
func main() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x) //o/p will vary due to race condition
}

//=================================

//solution using mutex

// var x = 0

// func increment(wg *sync.WaitGroup, m *sync.Mutex) {
// 	m.Lock() //only one Goroutine is allowed to execute this piece of code
// 	x = x + 1
// 	m.Unlock()
// 	wg.Done()
// }
// func main() {
// 	var w sync.WaitGroup
// 	var m sync.Mutex //Mutex is a struct type and we create a zero valued variable m of type Mutex
// 	for i := 0; i < 1000; i++ {
// 		w.Add(1)
// 		go increment(&w, &m) //If the mutex is passed by value instead of passing the address, each Goroutine will have its own copy of the mutex and the race condition will still occur.
// 	}
// 	w.Wait()
// 	fmt.Println("final value of x", x)
// }

//==========================================

//solve the race condition using channels

// var x = 0

// func increment(wg *sync.WaitGroup, ch chan bool) {
// 	ch <- true
// 	x = x + 1
// 	<-ch
// 	wg.Done()
// }
// func main() {
// 	var w sync.WaitGroup
// 	ch := make(chan bool, 1)
// 	for i := 0; i < 1000; i++ {
// 		w.Add(1)
// 		go increment(&w, ch)
// 	}
// 	w.Wait()
// 	fmt.Println("final value of x", x)
// }
