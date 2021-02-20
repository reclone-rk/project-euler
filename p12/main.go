// Highly divisible triangular number
package main

import (
	"fmt"
)

// Number of divisors
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

// Method -2
func numberOfDivisors(n int) int {
	divs := 1
	i := 1
	for i*i < n {
		if n%i == 0 {
			divs += 2
		}
		i++
	}
	if i*i == n {
		divs++
	}
	return divs
}

func solver(n int) int {
	arr := []int{}
	nth := 0
	sum := 0
	for len(arr) < 1000 {
		nth++
		sum += nth

		divs := numberOfDivisors(sum)

		for len(arr) <= divs {
			arr = append(arr, sum)
		}
	}
	return arr[n]
}

func main() {
	// What is the value of the first triangle number to have over five hundred divisors? (500)
	fmt.Println(solver(500))
	// fmt.Println(numberOfDivisors(28))
}
