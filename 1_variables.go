package main

// func main() {
// 	minInt(252)
// }

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	foo()
	fmt.Println("Go back to main")
	for i := 0; i < 100; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("Foo is Executed")
}
