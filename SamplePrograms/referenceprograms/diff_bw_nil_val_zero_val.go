// Go's _structs_ are typed collections of fields.
// They're useful for grouping data together to form
// records.

// The biggest problem in golang is to differentiate between actual zero value
// and the default zero value.
// Using pointers as struct fields helps to differentiate
// if there is no value the pointer would be nil
// else it would hold a value address
package referenceprograms

import "fmt"

type Foo struct {
	Present *bool `json:"foo"`
	Num     *int  `json:"number_of_foos"`
}

func sample() {
	var i int
	i = 2
	fm := Foo{Num: &i}
	fmt.Println(fm.Present, "value")
	var b bool
	b = false
	fm1 := Foo{Num: &i, Present: &b}
	fmt.Println(*(fm1.Present), "value")
}
