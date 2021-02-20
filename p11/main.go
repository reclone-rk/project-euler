package main

import (
	"bufio"
	"fmt"
	// "io/ioutil"
	// "log"
	"os"
)

func main() {
	// Text as a Whole
	// content, err := ioutil.ReadFile("input.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(content))

	// TEXT WORD BY WORD
	x := [20][20]int{}
	f, err := os.Open("input.txt")

    if err != nil {
        fmt.Println(err)
     }

    defer f.Close()

    scanner := bufio.NewScanner(f)
    scanner.Split(bufio.ScanWords)
	
	for i:=0; i<20; i++ {
		for j:=0; j<20; j++ {
			scanner.Scan()
			s := scanner.Text()
			var val int
			fmt.Sscanf(s, "%d", &val)
			// fmt.Println(val)
			x[i][j] = val
		}
	}
    // for scanner.Scan() {

    //     // fmt.Println(scanner.Text())
	// 	s := scanner.Text()
	// 	var val int
	// 	fmt.Sscanf(s, "%d", &val)
	// 	fmt.Println(val)
    // }
	// fmt.Println(x)
    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }

	var ans int
	for i:=0; i<20; i++ {
		for j:=0; j<20; j++ {
			
			down := 1
			for k:=0; k<4; k++ {
				xx := i + k
				yy := j
				if (xx >= 0 && xx < 20 && yy >=0 && yy<20) {
					down = down * x[xx][yy]
				}
			}
			if down > ans {
				ans = down
			}

			right := 1
			for k:=0; k<4; k++ {
				xx := i + k
				yy := j
				if (xx >= 0 && xx < 20 && yy >=0 && yy<20) {
					right = right * x[xx][yy]
				}
			}
			if right > ans {
				ans = right
			}

			diag1 := 1
			for k:=0; k<4; k++ {
				xx := i + k
				yy := j + k
				if (xx >= 0 && xx < 20 && yy >=0 && yy<20) {
					diag1 = diag1 * x[xx][yy]
				}
			}
			if diag1 > ans {
				ans = diag1
			}

			diag2 := 1
			for k:=0; k<4; k++ {
				xx := i - k
				yy := j + k
				if (xx >= 0 && xx < 20 && yy >=0 && yy<20) {
					diag2 = diag2 * x[xx][yy]
				}
			}
			if diag2 > ans {
				ans = diag2
			}
		}
	}
	fmt.Println(ans) // 70600674
}
