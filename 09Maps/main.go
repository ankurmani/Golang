package main

import "fmt"

func main() {
	hashmap := make(map[int]string)
	hashmap[0] = "ABCD"
	hashmap[1] = "CDEF"
	hashmap[2] = "EFGH"

	fmt.Println(hashmap)

	delete(hashmap, 1)

	fmt.Println(hashmap)

	for key, value := range hashmap {
		fmt.Printf("Key %v & value %v\n", key, value)
	}
}
