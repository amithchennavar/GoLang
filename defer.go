package main

import (
	"fmt"
)

// defer functions are basically execurted in a LIFO manner
func main8() {
	fmt.Println("main 1")
	defer fmt.Println(("defet 1"))
	fmt.Println("main 2")
	defer fmt.Println(("defet 2"))

}
