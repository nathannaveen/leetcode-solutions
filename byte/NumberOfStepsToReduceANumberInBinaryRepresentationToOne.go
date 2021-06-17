package main

func main() {

}

func numSteps(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '1' {
			res += 2
		} else {
			res++
		}
	}

	return res
}
