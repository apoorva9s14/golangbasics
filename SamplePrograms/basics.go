package main

import "fmt"

type MyInt int
type Speak interface {
	CanSpeak() string
}

func (j MyInt) CanSpeak() string {
	return "Speaking"
}
func main() {
	//Example of switch case
	switch i := 10; i {
	case 11:
		fmt.Println("I am 11")
	case 10:
		fmt.Println("I am 10")
		fallthrough
	default:
		fmt.Println("I am nobody")
	}
	a := make([]int, 0, 5)
	a = append(a, 1)
	fmt.Println(a)

	type Vertex struct {
		X int
		Y int
	}
	vertex_struct_var := Vertex{
		X: 1,
	}
	fmt.Println(vertex_struct_var)

	var s Speak = MyInt(0)
	fmt.Println(s.CanSpeak())
}
