package main

import "math"

func main() {

}

func minOperations(s string) int {
	h := make([]int, 2)
	for _, i := range s {
		if i == '0' {
			h[0]++
		} else {
			h[1]++
		}
	}
	return len(s)/2 - int(math.Min(float64(h[0]), float64(h[1])))
}
