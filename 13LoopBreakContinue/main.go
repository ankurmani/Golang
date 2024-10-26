package main

import "fmt"

func main() {
	i := func(a, b int) (int, int) { // anonymous function
		var i int
		for i = 0; i < 10; i++ {
			if i > 3 {
				break
			}
			fmt.Println(i)
		}
		return a, b
	}
	x, y := i(1, 2)
	fmt.Println(x, y)
	val := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(ProAdder(val...))
}

func ProAdder(values ...int) int { // variadic function, any number of argument
	total := 0

	for _, val := range values {
		total += val
	}

	return total
}
