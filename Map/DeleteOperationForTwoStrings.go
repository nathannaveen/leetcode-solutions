package main

import "math"

func main() {

}

func minDistance(word1 string, word2 string) int {
	first := make([]int, 26)
	second := make([]int, 26)
	res := 0

	for _, i := range word1 {
		first[i-'a']++
	}
	for _, i := range word2 {
		second[i-'a']++
	}

	for i := 0; i < 26; i++ {
		res += int(math.Abs(float64(first[i] - second[i])))
	}

	return res
}
