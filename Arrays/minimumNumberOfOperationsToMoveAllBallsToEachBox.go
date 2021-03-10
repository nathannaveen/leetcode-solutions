package main

func minOperationsBalls(boxes string) []int {
	res := make([]int, len(boxes))

	for i := 0; i < len(boxes); i++ {
		for j := 0; j < len(boxes); j++ {
			if boxes[j] == '1' && j != i {
				res[i] += abs(i - j)
			}
		}
	}
	return res
}
