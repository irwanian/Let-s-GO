package main

// func main() {
// 	minInt(252)
// }

import (
	"fmt"
)

func main() {
	var x = "I'm Hungry"
	fmt.Println(x)
	foo()
	fmt.Println("Go back to main")
	for i := 0; i < 7; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
	fmt.Println("Bella Ciao...")
}

func foo() {
	fmt.Println("Foo is Executed")
	for i := 1; i < 31; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 && i%3 != 0 {
			fmt.Println("Buzz")
		} else if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else {
			fmt.Println(i)
		}
	}
}
