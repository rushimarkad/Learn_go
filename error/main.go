// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {
// 	f, err := os.Open("/test.txt")
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Println(f.Name(), "opened successfully")
// }

//=============================================

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	f, err := os.Open("test.txt")
// 	if err != nil {
// 		var pErr *os.PathError
// 		if errors.As(err, &pErr) {
// 			fmt.Println("Failed to open file at path", pErr.Path)
// 			return
// 		}
// 		fmt.Println("Generic error", err)
// 		return
// 	}
// 	fmt.Println(f.Name(), "opened successfully")
// }

//================================================

//Custom errors
// var ErrBadPattern = errors.New("syntax error in pattern")

//===================================================

//custom errors

// package main

// import (
// 	"errors"
// 	"fmt"
// 	"math"
// )

// func circleArea(radius float64) (float64, error) {
// 	if radius < 0 {
// 		return 0, errors.New("Area calculation failed, radius is less than zero")
// 	}
// 	return math.Pi * radius * radius, nil
// }

// func main() {
// 	radius := -20.0
// 	area, err := circleArea(radius)
// 	if err != nil {
// 		fmt.Println(err)
// 		return
// 	}
// 	fmt.Printf("Area of circle %0.2f", area)
// }

//===================================================

//Additional details using errorf

package main

import (
	"fmt"
	"math"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Area of circle %0.2f", area)
}

//================================================
