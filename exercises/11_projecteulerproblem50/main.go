package main

import "fmt"

func main() {
	limit := 1000000
	var sum int
	x := 2
	isprime := func(i int) bool {
		var ans bool
		for x <= i {
			switch {
			case i%x == 0 && x < i:
				ans = false
				x = i + 1 //why does i++ here cause an error.
			case i%x != 0 && x < i:
				x++

			case i%x == 0 && x == i:

				ans = true

				x++

			}

		}

		return ans
	}
	var primelist []int
	if sum < limit {
		for i := 2; i < limit; i++ {
			if isprime(i) {
				sum += i
				if isprime(sum) && sum < limit {
					primelist = append(primelist, sum)

				}
			}
		}
	}
	fmt.Println(primelist[len(primelist)-1])

}

//The prime 41, can be written as the sum of six consecutive primes:

//41 = 2 + 3 + 5 + 7 + 11 + 13
//This is the longest sum of consecutive primes that adds to a prime below one-hundred.

//The longest sum of consecutive primes below one-thousand that adds to a prime, contains 21 terms, and is equal to 953.

//Which prime, below one-million, can be written as the sum of the most consecutive primes?
