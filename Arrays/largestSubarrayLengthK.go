package main

func main() {

}

func largestSubarray(nums []int, k int) []int {
	res := make([]int, k)
	for i := 0; i < len(nums)-k+1; i++ {
		res = whichArrayIsGreater(res, nums[i:i+k])
	}
	return res
}

func whichArrayIsGreater(one, two []int) []int {
	for i := 0; i < len(one); i++ {
		if one[i] > two[i] {
			return one
		} else if one[i] < two[i] {
			return two
		}
	}
	return []int{}
}
