package main

func main() {

}

func sortedSquares(nums []int) []int {
	left, right := 0, len(nums)-1
	res := []int{}

	for left <= right {
		l, r := nums[left]*nums[left], nums[right]*nums[right]
		if l > r {
			res, left = append([]int{l}, res...), left+1
		} else {
			res, right = append([]int{r}, res...), right-1
		}
	}

	return res
}
