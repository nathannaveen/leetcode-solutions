package main

func main() {

}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if absoluteValue(nums[i]-nums[j]) <= t && absoluteValue(i-j) <= k {
				continue
			}
			return false
		}
	}
	return true
}

func absoluteValue(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
