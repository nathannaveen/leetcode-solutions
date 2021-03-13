package main

func main() {
}

func subarraysDivByK(A []int, K int) int {
	res := 0

	for i := 1; i < len(A); i++ {
		for j := 0; j < len(A)-i+1; j++ {
			if sumArray(A[j:j+i])%K == 0 {
				res++
			}
		}
	}
	if sumArray(A)%K == 0 {
		res++
	}

	return res
}

func sumArray(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}
