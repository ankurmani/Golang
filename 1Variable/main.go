package main

import "fmt"

const Listout string = "ANKUR String"

func main() {
	var username string = "hitesh"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLog bool = true
	fmt.Println(isLog)
	fmt.Printf("Variable is of type: %T \n", isLog)

	var smallval uint8 = 255
	fmt.Println(smallval)
	fmt.Printf("Variable is of type: %T \n", smallval)

	var floatval float64 = 255.1111111111111111
	fmt.Println(floatval)
	fmt.Printf("Variable is of type: %T \n", floatval)

	// default values and some aliases

	var anotherval int
	fmt.Println(anotherval)
	fmt.Printf("Variable is of type: %T \n", anotherval)

	//implicit type

	var website = "learncodeline.in"
	fmt.Println(website)

	//no var style

	numberofUser := 30000
	fmt.Println(numberofUser)

	fmt.Println(Listout)
}
