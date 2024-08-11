package main

import (
	"fmt"

	"example/hello/data"
)

func init() {
	fmt.Println("Init1")
}

func init() {
	fmt.Println("Init2")
}

func main() {
	fmt.Println("Hello, World!")
	Function()
	var message string = "Hello, World!"
	message2 := "Hello, World!"
	message3 := "Hello, World!"
	fmt.Println(message, message2, message3)
	data.ExportThisFunction()
}
