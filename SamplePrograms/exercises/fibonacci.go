// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func fibonncci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonncci(n-2) + fibonncci(n-1)
	}
}

func mainFib() {

	fmt.Println(fibonncci(6))
}
