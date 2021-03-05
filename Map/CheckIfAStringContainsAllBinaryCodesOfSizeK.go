package main

func hasAllCodes(s string, k int) bool {
	m := make(map[string]int)

	for i := 0; i < len(s)-k+1; i++ {
		str := ""
		for j := i; j < i+k; j++ {
			str += string(s[j])
		}
		m[str]++
	}
	if k == 1 && len(m) == 2 {
		return true
	}
	return len(m) == k*k
}
