package main

// To print word count in a given string
import (
	"fmt"
	"strings"
)

func wordcount(s string) map[string]int {
	var m map[string]int
	m = make(map[string]int)
	for _, val := range strings.Fields(s) {
		// fmt.Println(val, ind)
		_, ok := m[val]
		// fmt.Println(each, ok)
		if ok {
			m[val] = m[val] + 1
		} else {
			m[val] = 1
		}

	}
	return m
}

func firstGoPgm() {
	// wc.Test(WordCount)
	fmt.Println(wordcount("A man a plan a canal panama A"))
}
