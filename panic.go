package main

import "fmt"

func main9() {
	dvd, dvsr := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dvd, dvsr, divide(dvd, dvsr))

	dvd, dvsr = 10, 0
	fmt.Printf("%v divided by %v is %v\n", dvd, dvsr, divide(dvd, dvsr))
}

func divide(dvd, dvsr int) int {

	defer func() {
		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}

	}()
	return dvd / dvsr
}
