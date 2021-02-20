package main

import "fmt"

func fixedSize2DArray(){
	x := [20][20]int{}
	for i:=0; i<20; i++ {
		for j:=0; j<20; j++ {
			x[i][j] = i + j
		}
	}
	fmt.Println(x)
}

func variableSize2DArray(){
	x := [][]int{}
	for i := 0; i<20; i++ {
		x = append(x, []int{})
		for j:=0; j<20; j++ {
			val := i + j
			x[i] =  append(x[i], val)
		}
	}
	fmt.Println(x)
}
func main(){
	fixedSize2DArray()
	variableSize2DArray()
}