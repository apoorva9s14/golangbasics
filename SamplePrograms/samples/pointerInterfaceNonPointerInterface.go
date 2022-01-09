// Sample of a pointer type may call the methods of its associated value type, but not vice versa
package samples

import "fmt"

type Mammal struct {
}
type samplePointerInterface interface {
	Speak() string
}

type valueInterface interface {
	Swim() string
}

func (m Mammal) Swim() string {
	return "Swimming!"
}

func (c *Mammal) Speak() string {
	return "Meow!"
}
func pointerInterface() {

	// Pointer interface can take only pointers
	var c samplePointerInterface
	// c = Mammal{} // This gives Error
	// cannot use (Mammal literal) (value of type Mammal) as samplePointerInterface value
	// in assignment: missing method Speak (Speak has pointer receiver)compilerInvalidIfaceAssign
	c = new(Mammal)
	fmt.Println(c.Speak())

	// value Interface can take both pointers and values
	var m valueInterface
	m = new(Mammal)
	m1 := Mammal{}
	fmt.Println(m.Swim())
	fmt.Println(m1.Swim())

}
