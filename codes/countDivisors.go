package main

import "fmt"

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

func main(){
	fmt.Println(numberOfDivisors(15))
}