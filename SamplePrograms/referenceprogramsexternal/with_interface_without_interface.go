package referenceprogramsexternal

import (
	"fmt"
	"reflect"
)

type Cat struct{}

func (c Cat) Say() string { return "meow" }

type Dog struct{}

func (d Dog) Say() string { return "woof" }

type Sayer interface {
	Say() string
}

func samplesWInterface() {
	c := Cat{}
	fmt.Println("Cat says:", c.Say())
	d := Dog{}
	fmt.Println("Dog says:", d.Say())
	animals := []Sayer{c, d}
	for _, a := range animals {
		fmt.Println(reflect.TypeOf(a).Name(), "says:", a.Say())
	}
	fmt.Println(1+2, "pop"+"pp")
}
