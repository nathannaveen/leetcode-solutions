package main

func findLeastNumOfUniqueInts(arr []int, k int) int {
	nums := [][]int{}
	numsCounter := 0

	for _, i := range arr {
		contain := false

		for j := range nums {
			if nums[j][0] == i {
				nums[j][1]++
				contain = true
			}
		}
		if !contain {
			nums = append(nums, []int{i, 1})
		}
	}

	for i := 1; i < len(nums); i++ {
		if i >= 1 && nums[i][1] < nums[i-1][1] {
			nums[i], nums[i-1] = nums[i-1], nums[i]
			i -= 2
		}
	}
	elements := len(nums)

	for i := 0; i < k; i++ {
		if nums[numsCounter][1] == 1 {
			numsCounter++
			elements -= 1
		} else {
			nums[numsCounter][1]--
		}

	}
	return elements
}
