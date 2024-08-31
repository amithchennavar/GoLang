package main

import (
	"fmt"
)

func main4() {
	var m map[string]int
	fmt.Println("bedfore: ", m)

	m = map[string]int{"foo": 1, "bar": 2}
	fmt.Println("After: ", m)

	fmt.Println(m["foo"])
	m["bar"] = 99 // update value to map

	delete(m, "foo") // delete
	m["baz"] = 418   // add value to map

	fmt.Println(m)

	fmt.Println(m["foo"])

}
