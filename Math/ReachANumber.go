package main

import "fmt"

func main() {
	fmt.Println(reachNumber(2))
}

func reachNumber(target int) int {
	if target < 0 {
		target = -target
	}
	counter := 1
	res := 0

	for res != target {
		if res+counter > target {
			res -= counter
		} else if res+counter < target {
			res += counter
		} else if res+counter == target {
			return counter
		}
		counter++
	}
	return counter
}
