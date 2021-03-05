package main

func main() {

}

func productExceptSelf(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		product := 1
		for j := 0; j < len(nums); j++ {
			if nums[i] != nums[j] {
				product *= nums[j]
			}
		}
		nums[i] = product
		if i != 0 {
			nums[i] /= nums[i-1]
		}
	}
	return nums
}
