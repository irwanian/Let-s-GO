package main

import (
	"fmt"
)

func main() {
	upperTriangle()
	bottomTriangle()
}

func upperTriangle() {
	output := ""
	inputValue := 5

	for i := 0; i < inputValue; i++ {
		for j := inputValue; j > i; j-- {
			output += "  "
		}
		for k := 0; k <= i; k++ {
			output += "* "
		}
		for l := 0; l < i; l++ {
			output += "* "
		}
		if i < (inputValue - 1) {
			output += "\n"
		}
	}
	fmt.Println(output)
}

func bottomTriangle() {
	output := ""
	inputValue := 5

	for i := 0; i <= inputValue; i++ {
		for j := 0; j < i; j++ {
			output += "  "
		}
		for k := inputValue; k >= i; k-- {
			output += "* "
		}
		for l := inputValue; l > i; l-- {
			output += "* "
		}
		if i < (inputValue) {
			output += "\n"
		}
	}
	fmt.Println(output)
}
