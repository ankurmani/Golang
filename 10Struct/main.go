package main

import "fmt"

type User struct {
	Name            string
	Email           string
	MobileNo        string
	NumberofFreinds int
}

func main() {
	userbook := User{"Ankur", "gmtripathi70@gmail.com", "7007403896", 16}

	fmt.Println(userbook)
	fmt.Printf("userbook data : %+v\n", userbook)
}
