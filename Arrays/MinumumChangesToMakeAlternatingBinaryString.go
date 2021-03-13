package main

func main() {

}

func minOperations(s string) int {

	onesAndZeros := make([]int, 2)
	for _, i := range s {
		if i == '1' {
			onesAndZeros[1]++
		} else {
			onesAndZeros[0]++
		}
	}

	if len(s)%2 == 0 {
		return max(onesAndZeros[0], onesAndZeros[1]) - len(s)/2
	} else {
		return len(s)/2 - min(onesAndZeros[0], onesAndZeros[1])
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
