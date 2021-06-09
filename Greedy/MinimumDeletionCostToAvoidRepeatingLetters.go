package main

import "fmt"

func main() {
	fmt.Println(minCost("aabaa", []int{1, 2, 3, 4, 1}))
}

func minCost(s string, cost []int) int {
	res := 0

	for i := 0; i < len(s)-1; i++ {
		fmt.Println(s)
		if s[i] == s[i+1] {
			if cost[i] < cost[i+1] {
				s = s[:i] + s[i+1:]
				res += cost[i]
			} else if i < len(s)-2 {
				s = s[:i+1] + s[i+2:]
				res += cost[i+1]
			}
			i -= 2
		}
	}
	if s[len(s)-2] == s[len(s)-2+1] {
		if cost[len(s)-2] < cost[len(s)-2+1] {
			s = s[:len(s)-2] + s[len(s)-2+1:]
			res += cost[len(s)-2]
		} else {
			s = s[:len(s)-2+1] + s[len(s)-2+2:]
			res += cost[len(s)-2+1]
		}
	}
	fmt.Println(s)
	return res
}
