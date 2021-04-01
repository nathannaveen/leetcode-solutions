package main

import (
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func main() {

}

func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	res := make([]bool, len(l))

	for i := 0; i < len(l); i++ {
		h := nums[l[i] : r[i]+1]
		for j := 1; j < len(h); j++ {
			if i >= 1 && h[i-1] > h[i] {
				h[i-1], h[i] = h[i], h[i-1]
				i -= 2
			}
		}
		fmt.Println(h)
		if len(h) <= 2 {
			res[i] = false
		} else {
			isArithmetic := true
			for j := 2; j < len(h)-1; j++ {
				if h[j+1]-h[j] != h[1]-h[0] {
					res[i] = false
					isArithmetic = false
					break
				}
			}
			if isArithmetic {
				res[i] = true
			}
		}

	}
	return res
}
