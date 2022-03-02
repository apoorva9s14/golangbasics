// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func mainFact() {

	fmt.Println(fact(4))
}
