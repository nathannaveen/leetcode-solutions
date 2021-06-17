package main

func main() {

}

func subarraysDivByK(nums []int, k int) int {
	res := 0
	for i := 0; i < len(nums); i++ {
		temp := nums[i]
		if temp%k == 0 {
			res++
		}
		for j := i + 1; j < len(nums); j++ {
			temp += nums[j]
			if temp%k == 0 {
				res++
			}
		}
	}
	return res
}
