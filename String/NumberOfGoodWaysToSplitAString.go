package main

func main() {
}

func numSplits(s string) int {
	res := 0

	for i := 1; i < len(s); i++ {
		first := make([]int, 26)
		firstCounter := 0
		second := make([]int, 26)
		secondCounter := 0

		for j := 0; j < i; j++ {
			if first[int(rune(s[j])-'a')] == 0 {
				firstCounter++
			}
			first[int(rune(s[j])-'a')]++
		}

		for j := i; j < len(s); j++ {
			if second[int(rune(s[j])-'a')] == 0 {
				secondCounter++
			}
			second[int(rune(s[j])-'a')]++
		}
		if firstCounter == secondCounter {
			res++
		}
	}

	return res
}
