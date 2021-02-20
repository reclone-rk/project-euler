	// Longest Collatz sequence
package main

import "fmt"

func counter(n int) int {
	if n == 1 {
		return 1
	}
	if n%2 == 0 {
		n = n / 2
	} else {
		n = 3*n + 1
	}
	return 1 + counter(n)
}

func main() {
	// 1 Million = 1,000,000
	var ans int
	var as int
	for i := 1; i < 1000000; i++ {
		val := counter(i)
		if val > ans {
			ans = counter(i)
			as = i
		}
	}
	fmt.Println(as)
}
