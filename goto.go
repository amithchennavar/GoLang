package main

import "fmt"

func main10() {
	fmt.Println("Start")

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j >= 10 {
				goto exit
			}
			fmt.Printf("i=%d, j=%d\n", i, j)
		}
	}
	fmt.Println("This will not be printed")

exit:
	fmt.Printf("Loop Exit")
}
