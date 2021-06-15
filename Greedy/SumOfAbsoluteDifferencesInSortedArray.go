package main

func main() {

}

func getSumAbsoluteDifferences(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			temp := abs(nums[i] - nums[j])
			res[i] += temp
			res[j] += temp
		}
	}
	return res
}
