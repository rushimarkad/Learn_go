package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
)

//==================================================

//reading from the file

// func main() {
// 	contents, err := os.ReadFile("test.txt") //returns a byte slice which is stored in contents
// 	if err != nil {
// 		fmt.Println("File reading error", err)
// 		return
// 	}
// 	fmt.Println("Contents of file:", string(contents)) //onvert contents to a string
// }

//================================================================

//Using absolute file path

// func main() {
// 	contents, err := os.ReadFile("/home/rushikesh/Learn_Go/filehandling/test.txt")
// 	if err != nil {
// 		fmt.Println("File reading error", err)
// 		return
// 	}
// 	fmt.Println("Contents of file:", string(contents))
// }

//====================================================

//Passing the file path as a command line flag

//we can get the file path as input argument from the command line and then read its contents.
//The flag package has a String function. This function accepts 3 arguments. The first is the name of the flag,
//second is the default value and the third is a short description of the flag.

// func main() {
// 	fptr := flag.String("fpath", "test.txt", "file path to read from")
// 	flag.Parse()
// 	contents, err := os.ReadFile(*fptr)
// 	if err != nil {
// 		fmt.Println("File reading error", err)
// 		return
// 	}
// 	fmt.Println("Contents of file:", string(contents))
// }

//==============================================================

// Bundling the text file along with the binary

//able to bundle the text file along with our binary
// After importing the embed package, the //go:embed directive can be used to read the contents of the file.

//#go:embed test.txt

// var contents []byte

// func main() {
// 	fmt.Println("Contents of file:", string(contents))
// }

//=================================================================

//Reading a file in small chunks

// func main() {
// 	fptr := flag.String("fpath", "test.txt", "file path to read from")
// 	flag.Parse()

// 	f, err := os.Open(*fptr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer func() {
// 		if err = f.Close(); err != nil {
// 			log.Fatal(err)
// 		}
// 	}()

// 	r := bufio.NewReader(f)
// 	b := make([]byte, 3) //read in the chunks of 3 bytes
// 	for {
// 		n, err := r.Read(b)
// 		if err == io.EOF {
// 			fmt.Println("finished reading file")
// 			break
// 		}
// 		if err != nil {
// 			fmt.Printf("Error %s reading file", err)
// 			break
// 		}
// 		fmt.Println(string(b[0:n]))
// 	}
// }

//========================================================

//Reading a file line by line

//following are the steps involved in reading a file line by line.
//Open the file
//Create a new scanner from the file
//Scan the file and read it line by line.

func main() {
	fptr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	f, err := os.Open(*fptr) //open the file using the path passed from the command line flag
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	s := bufio.NewScanner(f)
	for s.Scan() { //reads the next line of the file and the string that is read will be available through the text() method.
		fmt.Println(s.Text())
	}
	err = s.Err()
	if err != nil {
		log.Fatal(err)
	}
}
