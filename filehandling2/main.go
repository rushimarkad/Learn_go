package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"
)

//==================================================

// func main() {
// 	f, err := os.Create("test.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	l, err := f.WriteString("Hello World1")
// 	if err != nil {
// 		fmt.Println(err)
// 		f.Close()
// 		return
// 	}
// 	fmt.Println(l, "bytes written successfully")
// 	err = f.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

//=====================================================

//Writing bytes to a file

// func main() {
// 	f, err := os.Create("/home/rushikesh/Learn_Go/filehandling2/bytes")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	d2 := []byte{104, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100}
// 	n2, err := f.Write(d2)
// 	if err != nil {
// 		fmt.Println(err)
// 		f.Close()
// 		return
// 	}
// 	fmt.Println(n2, "bytes written successfully")
// 	err = f.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// }

//===========================================================

//Writing strings line by line to a file

// func main() {
// 	f, err := os.Create("lines.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		f.Close()
// 		return
// 	}
// 	d := []string{"Welcome to the world of Go1.", "Go is a compiled language.", "It is easy to learn Go."}

// 	for _, v := range d {
// 		fmt.Fprintln(f, v)
// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}
// 	}
// 	err = f.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("file written successfully")
// }

//========================================================

//Appending to a file

// func main() {
// 	f, err := os.OpenFile("lines.txt", os.O_APPEND|os.O_WRONLY, 0644) // open the file in append and write only mode
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	newLine := "File handling is easy."
// 	_, err = fmt.Fprintln(f, newLine) //Add new line to file
// 	if err != nil {
// 		fmt.Println(err)
// 		f.Close()
// 		return
// 	}
// 	err = f.Close()
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println("file appended successfully")
// }

//==================================================

//Writing to file concurrently

//When multiple goroutines write to a file concurrently, we will end up with a race condition.
//Hence concurrent writes to a file should be coordinated using a channel.

func produce(data chan int, wg *sync.WaitGroup) {
	n := rand.Intn(999)
	data <- n
	wg.Done()
}

func consume(data chan int, done chan bool) {
	f, err := os.Create("concurrent.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	for d := range data {
		_, err = fmt.Fprintln(f, d)
		if err != nil {
			fmt.Println(err)
			f.Close()
			done <- false
			return
		}
	}
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		done <- false
		return
	}
	done <- true
}

func main() {
	data := make(chan int)
	done := make(chan bool)
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(data, done)
	go func() {
		wg.Wait()
		close(data)
	}()
	d := <-done
	if d == true {
		fmt.Println("File written successfully")
	} else {
		fmt.Println("File writing failed")
	}
}
