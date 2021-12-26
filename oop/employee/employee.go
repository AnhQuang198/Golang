package employee

import "fmt"

type Employee struct {
	FirstName string
	LastName  string
}

func (e Employee) GetInfo() {
	fmt.Printf("First name %s - Last name %s", e.FirstName, e.LastName)
}
