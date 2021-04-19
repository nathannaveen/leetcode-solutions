package main

func main() {

}

func maxAscendingSum(nums []int) int {
	maximum := nums[0]
	sum := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[i-1] {
			sum += nums[i]
		} else {
			maximum, sum = max(maximum, sum), nums[i]
		}
	}
	maximum = max(maximum, sum)
	return maximum
}
