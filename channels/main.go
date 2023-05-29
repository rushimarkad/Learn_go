package main

import (
	"fmt"
	"sync"
	"time"
)

//==================================================

// func hello(done chan bool) {
// 	fmt.Println("hello go routine is going to sleep")
// 	time.Sleep(4 * time.Second)
// 	fmt.Println("hello go routine awake and going to write to done")
// 	done <- true
// }
// func main() {
// 	done := make(chan bool)
// 	fmt.Println("Main going to call hello go goroutine")
// 	go hello(done)
// 	<-done
// 	fmt.Println("Main received data")
// }

//===================================================

// func calcSquares(number int, squareop chan int) {
// 	sum := 0
// 	for number != 0 {
// 		digit := number % 10
// 		sum += digit * digit
// 		number /= 10
// 	}
// 	squareop <- sum
// }

// func calcCubes(number int, cubeop chan int) {
// 	sum := 0
// 	for number != 0 {
// 		digit := number % 10
// 		sum += digit * digit * digit
// 		number /= 10
// 	}
// 	cubeop <- sum
// }

// func main() {
// 	number := 589
// 	sqrch := make(chan int)
// 	cubech := make(chan int)
// 	go calcSquares(number, sqrch)
// 	go calcCubes(number, cubech)
// 	squares, cubes := <-sqrch, <-cubech
// 	fmt.Println("Final output", squares+cubes)
// }

//================================================

// func main() {
// 	ch := make(chan int)
// 	ch <- 5	//program will panic here
// }

//=================================================

// func sendData(sendch chan<- int) {
// 	sendch <- 10
// }

// func main() {
// 	sendch := make(chan<- int)	//create unidirectional channel
// 	go sendData(sendch)
// 	fmt.Println(<-sendch)
// }

//===================================================

//use bidirectional channel as unidirectional

// func sendData(sendch chan<- int) {	//converted to send only channel
// 	sendch <- 10
// }

// func main() {
// 	chnl := make(chan int)	//// bidirectional channel
// }
// 	go sendData(chnl)
// 	fmt.Println(<-chnl)

// channel is send only inside the sendData Goroutine but it's bidirectional in the main Goroutine.

//====================================================

//Closing channels

// func producer(chnl chan int) {
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// 	close(chnl) //close the channel
// }
// func main() {
// 	ch := make(chan int)
// 	go producer(ch)
// 	for {
// 		v, ok := <-ch //ok will be false if we are reading from closed channel
// 		if ok == false {
// 			break
// 		}
// 		fmt.Println("Received ", v, ok)
// 	}
// }

//==================================================

//using for range loop

// func producer(chnl chan int) {
// 	for i := 0; i < 10; i++ {
// 		chnl <- i
// 	}
// 	close(chnl)
// }
// func main() {
// 	ch := make(chan int)
// 	go producer(ch)
// 	for v := range ch {
// 		fmt.Println("Received ", v)
// 	}
// }

//=====================================

//Bufferred channel
//Pass additional parameter as capacity in make i.e ch := make(chan type, capacity)

// func main() {
// 	ch := make(chan string, 2) //does not block channel for writing upto 2 strings
// 	ch <- "naveen"
// 	ch <- "paul"
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

//============================================

//values to the channel are written in a concurrent Goroutine and read from the main Goroutine

// func write(ch chan int) {
// 	for i := 0; i < 5; i++ {
// 		ch <- i
// 		fmt.Println("successfully wrote", i, "to ch")
// 	}
// 	close(ch)
// }
// func main() {
// 	ch := make(chan int, 2)
// 	go write(ch)
// 	time.Sleep(2 * time.Second) //write 2 values to chan and waits for another goroutine to read from chan
// 	for v := range ch {
// 		fmt.Println("read value", v, "from ch")
// 		time.Sleep(2 * time.Second)

// 	}
// }

//=========================================================

//Deadlock

// func main() {
// 	ch := make(chan string, 2)
// 	ch <- "naveen"
// 	ch <- "paul"
// 	ch <- "steve"	//blocked here as capacity exceeds
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

//=====================================================

//Closing buffered channels
//It's possible to read data from a already closed buffered channel
//The channel will return the data that is already written to the channel and once all the data has been read, it will return the zero value of the channel.

// func main() {
// 	ch := make(chan int, 5)
// 	ch <- 5
// 	ch <- 6
// 	close(ch)
// 	n, open := <-ch
// 	fmt.Printf("Received: %d, open: %t\n", n, open)
// 	n, open = <-ch
// 	fmt.Printf("Received: %d, open: %t\n", n, open)
// 	n, open = <-ch
// 	fmt.Printf("Received: %d, open: %t\n", n, open)
// }

//==========================================================

//similar program using for range loop

// func main() {
// 	ch := make(chan int, 5)
// 	ch <- 5
// 	ch <- 6
// 	close(ch)
// 	for n := range ch {
// 		fmt.Println("Received:", n)
// 	}
// }

//================================================

//Length vs Capacity

//capacity of a buffered channel is the number of values that the channel can hold
//The length of the buffered channel is the number of elements currently queued in it.

// func main() {
// 	ch := make(chan string, 3)
// 	ch <- "naveen"
// 	ch <- "paul"
// 	fmt.Println("capacity is", cap(ch))
// 	fmt.Println("length is", len(ch))
// 	fmt.Println("read value", <-ch)
// 	fmt.Println("new length is", len(ch))
// }

//========================================================

//WaitGroup

//It is used to wait for a collection of Goroutines to finish executing

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done() //Decrements the counter by 1
}

func main() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)          //Increments the counter by 1
		go process(i, &wg) //If the pointer is not passed, then each Goroutine will have its own copy of the WaitGroup and main will not be notified when they finish executing.
	}
	wg.Wait() //wait method blocks the Goroutine in which it's called (here main go routine) until the counter becomes zero.
	fmt.Println("All go routines finished executing")
}
