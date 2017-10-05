package main

import "fmt"

func greatest(x ...int) int{
	var num int
	for _, v := range x {
		if v > num {
			num = v
		}
	}
return num
}

func main() {
	n := greatest(12, 13, 14, 15, 99, 102, 1000)
	fmt.Println(n)

}
