package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This needs to go in a file -- Linescode.in"

	file, err := os.Create("./mytext.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	error_msg(err)

	fmt.Println(length)

	defer file.Close()

	FileReader("./mytext.txt")
}

func FileReader(filename string) {
	fileByte, err := ioutil.ReadFile(filename)
	error_msg(err)
	fmt.Println(string(fileByte))
}

func error_msg(err error) {
	if err != nil {
		panic(err)
	}
}
