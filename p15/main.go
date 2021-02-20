// Lattice paths

package main

import (
	"fmt"
	"strconv"
)

var mp = map[string]int{}

func rec(i, j int) int {
	if i == 0 || j == 0 {
		return 1
	}
	s := strconv.Itoa(i) + "_" + strconv.Itoa(j)
	if mp[s] > 0 {
		return mp[s]
	}
	mp[s] = rec(i-1, j) + rec(i, j-1)
	return mp[s]
}

func main() {
	fmt.Println(rec(20, 20))
}
