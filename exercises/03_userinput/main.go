package main

import "fmt"

func main() {
	var input string
	fmt.Println("What is your name?")
	fmt.Scan(&input)
	fmt.Println("Your name is", input)
}
