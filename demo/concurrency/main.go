package main

import (
	"fmt"
	"os"
	"sync"
)

func writer(file *os.File, ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	lines := []string{"Line 1: This is line 1", "Line 2: This is line 2", "Line 3: This is line 3"}

	for _, line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			return
		}

		ch <- line // Send line to the reader
	}

	close(ch) // Close the channel to signal the reader that writing is done
}

func reader(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	// Read data from channel
	for line := range ch {
		fmt.Println("Data read from file:", line)
	}

}

func main() {
	fileName := "data.txt"

	// Create a file
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Failed to create file:", err)
		return
	}
	defer file.Close()

	// Create a channel to communicate between writer and reader
	ch := make(chan string)

	// Use a WaitGroup to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2) // Number of goroutines (writer and reader)

	// Concurrent writer and reader goroutine
	go writer(file, ch, &wg)
	go reader(ch, &wg)

	// Wait for all goroutines to finish
	wg.Wait()
}
