package main

import (
	"math"
)

func main() {
}

func findShortestSubArray(nums []int) int {
	h := make(map[int]int)
	max := []int{}
	n := []int{}
	for _, num := range nums {
		h[num]++
		if len(max) > 0 && max[len(max)-1] < h[num] {
			max = max[:]
			max = append(max, h[num])
			n = append(n, num)
		} else if len(max) == 0 {
			max = append(max, h[num])
		} else if max[len(max)-1] == h[num] {
			max = append(max, h[num])
		}
	}

	minCounter := 0
	for i := range max {
		counter := 0
		numberOfNums := 0
		start := false
		for _, num := range nums {
			if start == false && num == n[i] {
				start = true
				counter++
				numberOfNums++
			} else if start == true && num != n[i] {
				counter++
			} else if start == true && num == n[i] {
				numberOfNums++
				counter++
			}
			if numberOfNums == max[i] {
				break
			}
		}
		minCounter = int(math.Min(float64(counter), float64(minCounter)))
	}
	return minCounter
}
