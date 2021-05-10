package main

func main() {

}

func minOperations(nums []int) int {
	res := 0
	previous := nums[0]
	for i := 1; i < len(nums); i++ {
		res += max(0, previous+1-nums[i])
		previous = max(nums[i], nums[i-1]+1)
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
