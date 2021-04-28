package main

import "sort"

func main() {

}

func lastStoneWeight(stones []int) int {
	sort.Ints(stones)

	for len(stones) > 1 {
		if stones[0] == stones[1] {
			stones = stones[2:]
		} else if stones[0] > stones[1] {
			stones[0] = stones[0] - stones[1]
			stones = stones[1:]
		} else {
			stones[0] = stones[1] - stones[0]
			stones = stones[1:]
		}
	}

	if len(stones) == 1 {
		return stones[0]
	}
	return 0
}
