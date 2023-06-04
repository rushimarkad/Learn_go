package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Struct to store details of users
type details struct {
	name    string
	address string
	mobile  uint64
	pincode int32
}

var m = make(map[int]details)

var id = 0

// Check whether given id is present or not
func isIdPresent(id int) bool {
	_, ok := m[id]
	return ok
}

func addNew(id int) {
	var d details
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter name: ")
	scanner.Scan()
	d.name = scanner.Text()
	err1 := scanner.Err()
	if err1 != nil {
		log.Fatal(err1)
	}

	fmt.Print("Enter address: ")
	scanner.Scan()
	d.address = scanner.Text()
	err2 := scanner.Err()
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Print("Enter mobile number: ")
	fmt.Scanln(&d.mobile)

	fmt.Print("Enter pincode: ")
	fmt.Scanln(&d.pincode)

	m[id] = d
	fmt.Println("Details Saved with id =", id, "Successfully !!!\n")
}

func getDetails(id1 int) {
	ispresent := isIdPresent(id1)
	if ispresent {
		fmt.Println("Details corresponding to this id are:")
		fmt.Println("Name=", m[id1].name)
		fmt.Println("Address=", m[id1].address)
		fmt.Println("Mobile=", m[id1].mobile)
		fmt.Println("Pincode=", m[id1].pincode)
		fmt.Println("\n")
	} else {
		fmt.Println("Sorry...Id", id1, "does not exist !\n")
	}

}

func del(id1 int) {
	ispresent := isIdPresent(id1)
	if ispresent {
		delete(m, id1)
		fmt.Println("Details with id ", id1, "is deleted successfully")
	} else {
		fmt.Println("Cannot delete because Id", id1, "does not exist !\n")
	}

}

func show() {
	if len(m) == 0 {
		fmt.Println("Currently directory is empty !\n")
	} else {
		for key, value := range m {
			fmt.Println("id=", key)
			fmt.Println("Name=", value.name)
			fmt.Println("Address=", value.address)
			fmt.Println("Mobile=", value.mobile)
			fmt.Println("Pincode=", value.pincode)
		}
	}
}

func main() {
	var option int
	for {
		fmt.Println("\n=========Menu==========\n")
		fmt.Println("1. Get details with id")
		fmt.Println("2. Add new entry")
		fmt.Println("3. Delete from directory")
		fmt.Println("4. List from the directory\n")
		fmt.Print("Enter your choice: ")
		fmt.Scanln(&option)

		if option != 0 && option <= 4 {
			switch option {
			case 1:
				var id1 int
				fmt.Print("Enter id to get details: ")
				fmt.Scanln(&id1)
				getDetails(id1)
			case 2:
				id++
				addNew(id)
			case 3:
				var id1 int
				fmt.Print("Enter id to delete: ")
				fmt.Scanln(&id1)
				del(id1)
			case 4:
				show()
			}
		} else {
			fmt.Println("Invalid choice\n")
		}

		var ans string
		fmt.Print("Do you want to continue? (Y/N): ")
		fmt.Scanln(&ans)
		if ans == "N" || ans == "n" {
			break
		} else if ans == "Y" || ans == "y" {
			continue
		} else {
			fmt.Println("Invalid choice")
		}
	}
}
