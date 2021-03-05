package main

import "sort"

func main() {

}

func mctFromLeafValues(arr []int) int {
	sort.Ints(arr)
	sum := 0
	for i := 0; i < len(arr)-1; i++ {
		sum += arr[i] * arr[i+1]
	}
	return sum
}
