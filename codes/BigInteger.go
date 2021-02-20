package main

import (
	"fmt"
	"math/big"
)

func stringToBigInt(s string)*big.Int {
	n := new(big.Int)
	n, ok := n.SetString(s, 10)
    if !ok {
        fmt.Println("SetString: error")
    }
	return n
}

func main(){
	sum := new(big.Int)
	s := "42846280183517070527831839425882145521227251250327"
	b := new(big.Int)
	b = stringToBigInt(s)
	fmt.Println(b)
	sum.Add(sum, b)
	fmt.Println(sum)
}