package main

func main() {

}

func numberOfSubarrays(nums []int, k int) int {
	res := 0
	oddCounter := 0
	start := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			oddCounter++
		}
		if oddCounter == k {
			start = i + 1
			res++
			break
		}
	}

	for i := start; i < len(nums); i++ {

		if nums[i]%2 == 1 {
			oddCounter++
			if nums[i-k]%2 == 1 {
				oddCounter--
			}
			if oddCounter == k {
				res++
			}
		}

	}
	return res
}
