package department

import "fmt"

type department struct {
	name string
}

func New(name string) department {
	d := department{
		name: name,
	}
	return d
}

func (d department) ShowInfo() {
	fmt.Printf("Department %s", d.name)
}
