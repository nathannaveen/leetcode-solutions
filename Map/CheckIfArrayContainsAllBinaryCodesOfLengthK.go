package main

func main() {
}

func hasAllCodes(s string, k int) bool {
	m := make(map[string]int)

	for i := 0; i < len(s)-k+1; i++ {
		m[s[i:i+k]]++
	}

	return len(m) == twoPow(k)
}

func twoPow(k int) int {
	res := 1
	for i := 0; i < k; i++ {
		res *= 2
	}
	return res
}
