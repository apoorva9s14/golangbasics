package main

import (
	"fmt"
	"math/rand"
)

var my_array []int

func main() {
	i := 0
	for i < 100 {
		my_array = append(my_array, rand.Intn(100))
		i += 1
	}
	even_count := 0
	odd_count := 0
	for i := 0; i < len(my_array); i++ {

		if my_array[i]%2 == 0 {

			even_count += 1

		} else {

			odd_count += 1

		}
	}
	fmt.Println(even_count, odd_count)
}
