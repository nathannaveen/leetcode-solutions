package main

func main() {

}

func countBinarySubstrings(s string) int {
	start := 0
	res := 0
	counter := 0

	for i := 0; i < len(s); i++ {
		if i != start+1 && counter == 0 {
			res++
			i = start + 1
			start++
		}
		if s[i] == '1' {
			counter++
		} else {
			counter--
		}
	}
	return res
}
