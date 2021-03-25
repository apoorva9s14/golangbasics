package main

import (
	"fmt"
	"strings"
)

func main() {
	var inp_str = "Culpa pariatur commodo cillum nulla duis veniam esse incididunt officia ex deserunt Occaecat sunt do elit dolor mollit Ipsum ullamco cupidatat amet nostrud enim excepteur duis do ex quis labore ut dolore Excepteur ut nostrud elit ut consectetur Elit adipisicing Lorem sit amet sit reprehenderit Lorem aliquip dolore id commodo ut Id quis irure ut eiusmod aliqua labore laborum amet magna aliqua mollit mollit laboris"
	var myarray = strings.Split(inp_str, " ")
	fmt.Printf("%T", myarray[0])
	word_count_map := make(map[int]int)
	for i := 0; i < len(myarray); i++ {
		wordCount := len(myarray[i])
		_, ok := word_count_map[wordCount]
		if ok {
			word_count_map[wordCount]++
		} else {
			word_count_map[wordCount] = 1
		}

	}
	fmt.Println(word_count_map)
	// fmt.Printf("%d\n", strings.Split(inp_str, "")[0].count())
	// fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))
	// fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	// fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))
}
