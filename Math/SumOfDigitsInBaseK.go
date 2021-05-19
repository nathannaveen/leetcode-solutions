package main

import "fmt"

func main() {
	fmt.Println(sumBase(68, 2))
}

func sumBase(n int, k int) int {
	res := 0
	for n > 0 {
		res += n % k
		n /= k
	}
	return res
}
