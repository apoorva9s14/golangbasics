package main

import "fmt"

type sampleInterface interface {
	add(y int) int
	concatenate(y string) string
}

// here both the samples implement the interface
// based on the value of int passed, we decide with sample struct to use
type sample struct {
	x int
	y string
}

func (s sample) add(y int) int {
	return s.x + y
}

func (s sample) concatenate(y string) string {
	return s.y + y
}

type sample1 struct {
	x int
	y string
}

func (s sample1) add(y int) int {
	return s.x + y
}

func (s sample1) concatenate(y string) string {
	return s.y + y
}

// the return value for returnImplementedType is sampleInterface,
// either of the struct samples can fit in that value
func returnImplementedType(x int) sampleInterface {
	if x == 0 {
		return &sample{
			x: 0,
			y: "pop",
		}
	} else if x == 1 {
		return &sample1{
			x: 1,
			y: "pop1",
		}
	}
	return nil
}
func main() {
	si := returnImplementedType(0)
	fmt.Println(si.add(1))
	fmt.Println(si.concatenate(" appended"))

	si = returnImplementedType(1)
	fmt.Println(si.add(1))
	fmt.Println(si.concatenate(" appended"))
}
