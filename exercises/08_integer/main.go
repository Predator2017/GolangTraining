package main

import "fmt"

func main() {

	num := 2
	handler := func(num int) (int, bool) {
		var answer bool
		if num%2 == 0 {
			answer = true
		}
		return num / 2, answer
	}
	fmt.Println(handler(num))
}
