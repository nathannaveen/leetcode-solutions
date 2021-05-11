package main

import "fmt"

func main() {
	fmt.Println(checkInclusion("hello", "ooolleoooleh"))
}

func checkInclusion(s1 string, s2 string) bool {
	h := make([]int, 26)
	g := make([]bool, 26)
	k := len(s1)
	counter := 0

	for _, i := range s1 {
		h[i-'a']++
		g[i-'a'] = true
	}

	for i := 0; i < k; i++ {
		if h[s2[i]-'a'] >= 1 {
			counter++
		}
	}

	for i := k; i < len(s2); i++ {
		if g[s2[i-k]-'a'] == true {
			h[s2[i-k]-'a']++
			counter--
		}
		if h[s2[i-k]-'a'] >= 1 {
			counter++
			h[s2[i-k]-'a']--
		}
		if counter == k {
			return true
		}
	}

	return false
}
