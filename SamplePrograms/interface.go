package main

import (
	"fmt"
)

func printInterface() {
	m := make(map[string]interface{})
	m["abc"] = "xyz"
	fmt.Println(m, m["abc"].(string))

}
