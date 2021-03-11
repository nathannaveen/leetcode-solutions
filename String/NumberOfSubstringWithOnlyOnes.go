package main

func main() {

}

func numSub(s string) int {
	res := 0

	if onlyOnes(s) {
		sum := 0
		for i := 1; i <= len(s); i++ {
			sum += i
		}
		return sum
	}

	for i := 1; i < len(s); i++ {
		for j := 0; j <= len(s)-i; j++ {
			if onlyOnes(s[j : j+i]) {
				res++
			}
		}
	}

	return res
}

func onlyOnes(s string) bool {
	for _, i := range s {
		if i != '1' {
			return false
		}
	}
	return true
}
