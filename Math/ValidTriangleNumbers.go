package main

func main() {

}

func triangleNumber(nums []int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i] > nums[j] && nums[i] > nums[k] && nums[i] < nums[j]+nums[k] {
					res++
				} else if nums[i] < nums[j] && nums[j] > nums[k] && nums[j] < nums[k]+nums[i] {
					res++
				} else if nums[k] > nums[j] && nums[i] < nums[k] && nums[i]+nums[j] > nums[k] {
					res++
				}
			}
		}
	}
	return res
}
