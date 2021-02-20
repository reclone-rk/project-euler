package main

import (
	"bufio"
	"fmt"
	"log"
	"math/big"
	"os"
	// "strconv"
)

func demoStringToInt64(s string) *big.Int {
	// var s string = "9223372036854775807"
	// fmt.Println(s)
	// a := big.NewInt(0)
	// i, _ := strconv.ParseInt(s, 10, 64)
	// fmt.Printf("val: %v ; type: %[1]T\n", i)
	// fmt.Println(i)
	n := new(big.Int)
    n, ok := n.SetString(s, 10)
    if !ok {
        fmt.Println("SetString: error")
        // return
    }
    // fmt.Println(n)
	return n
}

func main() {
	f, err := os.Open("input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	sum := new(big.Int)
	for scanner.Scan() {
		s := scanner.Text()
		val := new(big.Int)
		val = demoStringToInt64(s)
		// fmt.Sscanf(s, "%v", &val)
		sum.Add(sum, val)
		// fmt.Println(val)
		// fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sum)
	// demoStringToInt64()
}
