package main

import "fmt"

func main() {
	var fruitlist [4]string
	fruitlist[0] = "Apple"
	fruitlist[1] = "Tomato"
	fruitlist[2] = "Peach"
	fmt.Println("fruits are : ", fruitlist)

	var vegList = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("Veggies list size is ", len(vegList))
	fmt.Println("Veggies is ", vegList)
}
