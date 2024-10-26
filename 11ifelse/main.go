package main

import "fmt"

func main() {
	logincount := 23
	var result string

	if logincount < 10 {
		result = "i am at 10"
	} else if logincount > 10 && logincount < 20 {
		result = "i am between 10 and 20"
	} else {
		result = "i am above 20"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is more than 10")
	}

}
