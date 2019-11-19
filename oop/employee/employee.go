package main

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
	fmt.Printf("%s %s has %d leaves remaining", e.FirstName, e.LastName, (e.TotalLeaves - e.LeavesTaken))
}
func main() {
	e := Employee{
		FirstName:   "ab",
		LastName:    "hi",
		TotalLeaves: 22,
		LeavesTaken: 10,
	}
	e.LeavesRemaining()
}
