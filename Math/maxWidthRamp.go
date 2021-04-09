package main

func main() {

}

func maxWidthRamp(A []int) int {
	maximum := 0
	for j := len(A) - 1; j >= len(A)/2; j-- {
		for i := 0; i < len(A); i++ {
			if A[j] <= A[i] && j-i > maximum {
				maximum = j - i
			}
		}
	}
	return maximum
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
