package main

import "fmt"

func main() {
	intSlices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(intSlices)

	for _, each := range intSlices {
		if each%2 == 0 {
			fmt.Println(each, " is even")
		} else {
			fmt.Println(each, " is odd")
		}
	}
}
