package main

func main() {

}

func numberOfSubstrings(s string) int {
	res := 0
	for i := 3; i < len(s); i++ {
		for j := 0; j < len(s)-i+1; j++ {
			if containsAllThree(s[j : j+i]) {
				res++
			}
		}
	}
	if containsAllThree(s) {
		res++
	}
	return res
}

func containsAllThree(s string) bool {
	a, b, c := 0, 0, 0

	for _, i := range s {
		switch i {
		case 'a':
			a++
		case 'b':
			b++
		case 'c':
			c++
		}
		if a >= 1 && b >= 1 && c >= 1 {
			return true
		}
	}
	return false
}
