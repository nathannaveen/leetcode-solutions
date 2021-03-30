package main

func main() {

}

func minSubsequence(nums []int) []int {
	minimum := []int{}
	minSum := 0
	sum := 0
	for _, num := range nums {
		sum += num
	}

	arr := []int{}
	h := 0
	for i := 0; i < len(nums); i++ {
		if h+nums[i] > sum-(h+nums[i]) {
			if len(minimum) == 0 {
				minimum = arr
				minSum = h
			}
			if minSum > h && len(arr) < len(minimum) {
				minimum = arr
				minSum = h
			}
			h -= nums[len(nums)-(len(arr)-1)]
			arr = arr[1:]
		} else {
			arr = append(arr, nums[i])
			h += nums[i]
		}
	}

	return minimum
}
