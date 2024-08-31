package main

import (
	"fmt"
	"slices"
)

func main2() {

	arr := []string{"a", "b", "c"}
	fmt.Println(arr)

	arr = append(arr, "5")
	fmt.Println(arr)

	arr = slices.Delete(arr, 2, 3)
	fmt.Println(arr)
}
