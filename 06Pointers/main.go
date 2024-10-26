package main

import "fmt"

func main() {
	fmt.Println("Pointer implementation :")
	numbers := 23
	var ptr *int = &numbers
	fmt.Println("number ptr is ", ptr)
	fmt.Println("number is ", *ptr)

}
