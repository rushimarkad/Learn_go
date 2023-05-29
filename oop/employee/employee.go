package employee

import (
	"fmt"
)

type Employee struct {
	FirstName   string
	LastName    string
	TotalLeaves int
	LeavesTaken int
}

func (e Employee) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining\n", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}

// type employee struct {
//     firstName   string
//     lastName    string
//     totalLeaves int
//     leavesTaken int
// }

// func New(firstName string, lastName string, totalLeave int, leavesTaken int) employee {
//     e := employee {firstName, lastName, totalLeave, leavesTaken}
//     return e
// }

// func (e employee) LeavesRemaining() {
//     fmt.Printf("%s %s has %d leaves remaining\n", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
// }
