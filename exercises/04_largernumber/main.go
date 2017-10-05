package main

import "fmt"

func main() {
	var input1 int32
	var input2 int32
	fmt.Println("Please input a number:")
	fmt.Scan(&input1)
	fmt.Println("Please input a larger number:")
	fmt.Scan(&input2)
	var result int32
	result = input2 % input1
	fmt.Println("The remainder is:", result)

}
