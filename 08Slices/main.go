package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")
	var fruitlist = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("slice of fruitlist is %T\n", fruitlist)
	fruitlist = append(fruitlist, "Mango", "Banana")

	fmt.Println(fruitlist)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867

	highScores = append(highScores, 555, 666, 32)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)

	highScores = append(highScores[:3], highScores[4:]...)

	fmt.Println(highScores)
}
