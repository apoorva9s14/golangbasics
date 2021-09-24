package main

import "fmt"

type Duck interface {
	Quack()
}

type Donald struct {
}

func (d Donald) Quack() {
	fmt.Println("quack quack!")
}

type Daisy struct {
}

func (d Daisy) Quack() {
	fmt.Println("~quack ~quack")
}

type Dog struct {
}

func (d Dog) Bark() {
	fmt.Print("woof woof")
}

func sayQuack(duck Duck) {
	duck.Quack()
}

func interfaceSample1() {
	donald := Donald{}
	sayQuack(donald) // quack quack!

	daisy := Daisy{}
	sayQuack(daisy) // ~quack ~quack

	// This piece of code throws compile error
	// dog := Dog{}
	// sayQuack(dog) // compile error
}
