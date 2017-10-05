package main

import "fmt"

var sum int

func main() {

	for i := 1; i <= 1000; i++ {
		if i%3 == 0 {
			add(i)
		} else if i%5 == 0 {
			add(i)
		}
	}
	fmt.Println(sum)

}
func add(num int) {

	sum = sum + num

}
