package main

import "sort"

func main() {

}

func connectSticks(sticks []int) int {
	sum := 0
	sort.Ints(sticks)

	for len(sticks) != 1 {
		sum += sticks[0] + sticks[1]
		sticks[1] += sticks[0]
		sticks = sticks[1:]
		sort.Ints(sticks)
	}

	return sum
}
