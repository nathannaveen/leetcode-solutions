package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(isPowerOfFour(4))
}

func isPowerOfFour(n int) bool {

	for i := 0; i*i < n; i++ {
		if int(math.Pow(float64(4), float64(i))) == n {
			return true
		}
	}

	return false
}
