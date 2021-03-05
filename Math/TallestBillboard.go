package main

func tallestBillboard(rods []int) int {
	res, left, right := 0, 0, len(rods)-1
	leftSum, rightSum := rods[left], rods[right]
	for left < right {
		if leftSum < rightSum {
			left++
			leftSum += rods[left]
		} else if leftSum > rightSum {
			right--
			rightSum += rods[right]
		} else {
			res = max(res, rightSum)
			// we could also do max using leftSum, but it doesn't matter because they are equal
			left++
			right--
			rightSum += rods[right]
			leftSum += rods[left]
		}
	}
	if leftSum == rightSum {
		res = max(res, rightSum)
	}
	return res
}
