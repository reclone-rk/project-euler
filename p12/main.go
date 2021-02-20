// Highly divisible triangular number
package main

import (
	"fmt"
)

func divCount(n int) int {
	var hash = make([]bool, n+1)
	for i := 0; i <= n; i++ {
		hash[i] = true
	}

	for p := 2; p*p < n; p++ {
		if hash[p] == true {
			for i := p * 2; i < n; i += p {
				hash[i] = false
			}
		}
	}

	total := 1
	for p := 2; p <= n; p++ {
		if hash[p] {
			count := 0
			if n%p == 0 {
				for n%p == 0 {
					n = n / p
					count++
				}
				total = total * (count + 1)
			}
		}
	}
	return total
}

func nod(n int) int {
	var cnt int
	var i int
	for i = 1; i <= n; i++ {
		if n%i == 0 {
			cnt++
		}
	}
	return cnt
}

func main() {
	// What is the value of the first triangle number to have over five hundred divisors? (500)
	var nth int
	var sum int
	for {
		nth++
		sum += nth
		if divCount(sum) > 5 {
			fmt.Println(sum)
			break
		}
	}
}
