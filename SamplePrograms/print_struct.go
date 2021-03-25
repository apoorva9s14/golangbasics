package main

import "fmt"

type Name struct {
	name string
	id   int
}

type Emp struct {
	Name
	id     int
	status bool
}

func (n *Name) display() {

	n.id = 11
	fmt.Println("ID is changed to ", n.id)
}

func toggleemployeestatus(emp *Emp) {
	emp.status = !emp.status
	fmt.Println(emp.status)
}
func main() {
	myname := Name{}
	myname.id = 10
	myemp := new(Emp)
	// myemp := Emp{myname, 100}
	myemp.id = 0
	fmt.Printf("myemp:%T, myname:%T", myemp, myname)
	myemp.name = myname.name
	myemp.status = true
	toggleemployeestatus(myemp)
	fmt.Printf("myname_id:%v, myemp_id:%v", myname.id, myemp.id)
	fmt.Println(myemp.status)
	myname.display()
	fmt.Printf("myname_id:%v, myemp_id:%v", myname.id, myemp.id)

}
