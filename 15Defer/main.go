package main

import "fmt"

func main() {
	defer fmt.Println("Output A")
	defer fmt.Println("Output B")
	defer fmt.Println("Output C")
	fmt.Println("Output Char")

	defer_func()
}

// Output Char, 43210, Output C, Output B, Output A

func defer_func() {
	var i int
	for i = 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
