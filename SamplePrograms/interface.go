package main

import (
	"fmt"
	"time"
)

func printInterface() {
	s := 1 * time.Microsecond
	fmt.Println(s)
	m := make(map[string]interface{})
	m["abc"] = "xyz"
	fmt.Println(m, m["abc"].(string))

}
