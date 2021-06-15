package main

import "fmt"

func main() {
	fmt.Println(missingNumber([]int{100, 300, 400}))
}

func missingNumber(arr []int) int {
	diff1 := arr[1] - arr[0]
	diff2 := arr[2] - arr[1]
	if len(arr) == 3 {
		i, done := smallerThan3(arr, diff1, diff2, -1, diff1, diff2)
		if done {
			return i
		}
		i, done = smallerThan3(arr, diff1, diff2, 1, diff2, diff1)
		if done {
			return i
		}
	}
	diff3 := arr[3] - arr[2]
	diff := 0

	if diff1 == diff2 {
		diff = diff1
	} else if diff2 == diff3 {
		diff = diff2
	} else if diff3 == diff1 {
		diff = diff3
	}

	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] != diff {
			return arr[i] + diff
		}
	}
	return arr[0]
}

func smallerThan3(arr []int, diff1 int, diff2 int, neg int, diff12 int, diff22 int) (int, bool) {
	if diff1 > diff2 && diff2*neg < 0 {
		return arr[1] + diff12, true
	} else if diff2 < 0 {
		return arr[0] + diff22, true
	}
	return 0, false
}
