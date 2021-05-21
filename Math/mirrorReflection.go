package main

import "fmt"

func main() {
	fmt.Println(mirrorReflection(3, 2))
}

func mirrorReflection(p int, q int) int {
	for i := 2; i < max(p, q); i++ {
		if p%i == 0 && p != i {
			p /= i
		}
		if q%i == 0 && q != i {
			q /= i
		}
	}
	fmt.Println(p, q)
	if q%2 == 1 && p%2 == 0 || ((p/q)%2 == 0 && p%q == 0) {
		return 2
	}
	if q%2 == 1 && p%2 == 1 || ((p/q)%2 == 1 && p%q == 0) {
		return 1
	}
	if q%2 == 0 && p%2 == 1 {
		return 0
	}
	return -1
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
