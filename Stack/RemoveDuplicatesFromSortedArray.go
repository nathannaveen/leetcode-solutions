package main

func removeDuplicates(nums []int) int {
	h := []int{}
	for _, num := range nums {
		if len(h) > 2 {
			if h[len(h)-1] == num && h[len(h)-2] == num {
				continue
			} else {
				h = append(h, num)
			}
		} else {
			h = append(h, num)
		}
	}

	for i, i2 := range h {
		nums[i] = i2
	}

	return len(h)
}
