//Select is used to choose from multiple send/receive channel operations
//The select statement blocks until one of the send/receive operations is ready.
//If multiple operations are ready, one of them is chosen at random

package main

import (
	"fmt"
	"time"
)

//=======================================

// func server1(ch chan string) {
// 	time.Sleep(6 * time.Second)
// 	ch <- "from server1"
// }
// func server2(ch chan string) {
// 	time.Sleep(3 * time.Second)
// 	ch <- "from server2"

// }
// func main() {
// 	output1 := make(chan string)
// 	output2 := make(chan string)
// 	go server1(output1)
// 	go server2(output2)
// 	select {
// 	case s1 := <-output1:
// 		fmt.Println(s1)
// 	case s2 := <-output2:
// 		fmt.Println(s2)
// 	}
// }

//=============================

//Default case
//The default case in a select statement is executed when none of the other cases is ready.
//This is generally used to prevent the select statement from blocking.

// func process(ch chan string) {
// 	time.Sleep(10500 * time.Millisecond)
// 	ch <- "process successful"
// }

// func main() {
// 	ch := make(chan string)
// 	go process(ch)
// 	for {
// 		time.Sleep(1000 * time.Millisecond)
// 		select {
// 		case v := <-ch:
// 			fmt.Println("received value: ", v)
// 			return
// 		default:
// 			fmt.Println("no value received")
// 		}
// 	}

// }

//===================================================

//Deadlock and default case

// func main() {
// 	ch := make(chan string)
// 	select {
// 	case <-ch: //panics
// 	}
// }

//==================================================

//above program with default

// func main() {
// 	var ch chan string
// 	select {
// 	case v := <-ch:
// 		fmt.Println("received value", v)
// 	default:
// 		fmt.Println("default case executed")

// 	}
// }

//=====================================================

//Random selection
//When multiple cases in a select statement are ready, one of them will be executed at random.

func server1(ch chan string) {
	ch <- "from server1"
}
func server2(ch chan string) {
	ch <- "from server2"

}
func main() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	time.Sleep(1 * time.Second)
	select { //output will vary
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

//========================================

//Empty select

// func main() {
// 	select {} //it will block forever resulting in a deadlock
// }
