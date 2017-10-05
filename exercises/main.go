package main

import "fmt"

func greatest(x ...int) int {
	length := len(x)
	var num int
	for i := 0; i < length; i++ {
		if x[i] > num {
			num = x[i]
		}
	}
	return num
}

func main() {
	list := greatest(12, 13, 15, 18, 19, 25, 1000)

	fmt.Println(list)

}
