package main

import "fmt"

type Identity struct {
	Id   int
	Name string
}

func (i *Identity) PrintIdentity() {
	fmt.Printf("Id = %d, Name = %s\n", i.Id, i.Name)
}

type Employee struct {
	Identity
	IsEmployed bool
	City       string
}

func (e *Employee) ToggleEmploymentStatus() {
	e.IsEmployed = !e.IsEmployed
	fmt.Printf(" Inside toggle fn : %v\n", e)
}

func (e *Employee) Display() {
	fmt.Println(e)
}

type Department struct {
	Identity
}

func (d *Department) Display() {
	fmt.Println(d)
}

/* func (d *Department) PrintIdentity() {
	fmt.Printf("Id = %d, Name = %s\n", d.Id, d.Name)
} */
func main() {
	emp := Employee{Identity{100, "xyz"}, true, "Banglaore"}
	emp.PrintIdentity()
}
