package main

func main() {

}

func validSubarrays(nums []int) int {
	res := 0
	for i := 1; i < len(nums); i++ {
		for j := 0; j < len(nums)-i+1; j++ {
			// [j : j + i]
			res += firstIsSmallest(nums[j : j+i])
		}
	}
	res += firstIsSmallest(nums)
	return res
}

func firstIsSmallest(h []int) int {
	for i := 1; i < len(h); i++ {
		if h[0] > h[i] {
			return 0
		}
	}
	return 1
}
