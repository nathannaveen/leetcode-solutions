package main

func main() {

}

func minOperations(nums []int) int {
	res := 0
	for i := 1; i < len(nums); i++ {
		res += max(0, nums[i-1]+1-nums[i])
		nums[i] = max(nums[i], nums[i-1]+1)
	}
	return res
}
