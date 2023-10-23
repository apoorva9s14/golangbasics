// Method = Function + receiver
// here receiver is int1

package main

import "fmt"

type int1 int

var x int1 = 1

// method with pointer receiver
// pointer receivers are helpful when you want to modify the receiver
func (i *int1) myint() bool {
	fmt.Println("in myint")
	*i = *i + 1
	// modification in i effects x
	return true

}

// method with value receiver
func (i int1) myintVal() bool {
	fmt.Println("in myintVal")
	i = i + 1
	// modification in i will not affect x
	return true

}

// func main() {
// 	fmt.Println("Hello, 世界")

// 	fmt.Println(x.myint())
// 	fmt.Println(x)
// 	fmt.Println(x.myintVal())
// 	fmt.Println(x)
// }
