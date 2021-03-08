package main

func main() {

}

func waysToMakeFair(nums []int) int {
	res := 0

	for i := 0; i < len(nums); i++ {
		evenSum, oddSum := 0, 0
		hasRemoved := false
		for j := 0; j < len(nums); j++ {
			if j == i {
				hasRemoved = true
			} else if hasRemoved {
				if j%2 == 0 {
					oddSum += nums[j]
				} else {
					evenSum += nums[j]
				}
			} else {
				if j%2 == 0 {
					evenSum += nums[j]
				} else {
					oddSum += nums[j]
				}
			}
		}
		if oddSum == evenSum {
			res++
		}

	}

	return res
}
