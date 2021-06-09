package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(getSmallestString(5, 73))
}

func getSmallestString(n int, k int) string {
	k -= n
	res := make([]string, n)
	//
	//for i := 0; i < n; i++ {
	//	res[i] = "a"
	//}

	for i := len(res) - 1; i >= 0; i-- {
		fmt.Println(res)
		if k == 0 {
			res[i] = "a"
		} else if k <= 25 {
			res[i] = string(k + 'a')
			k = 0
		} else {
			res[i] = "z"
			k -= 25
		}
	}

	return strings.Join(res, "")
}
