package main

import (
	"bufio"
	"demo/menu"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)

func mainf() {
loop:
	for {
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")
		choice, _ := in.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch strings.TrimSpace(choice) {
		case "1":
			menu.PrintMenu()
		case "2":
			menu.AddItem()
		case "q":
			break loop
		default:
			fmt.Println("Please choose withing given values")

		}
	}

}
