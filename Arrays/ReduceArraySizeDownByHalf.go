package main

import (
	"sort"

	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func main() {

}

func minSetSize(arr []int) int {
	sort.Ints(arr)
	size := []int{}
	num := arr[0]
	counter := 0

	for _, i := range arr {
		if num == i {
			counter++
		} else {
			num = i
			size = append(size, counter)
			counter = 1
		}
	}

	res := 0
	g := len(size) - 1
	sort.Ints(size)
	left, right := 0, len(size)-1
	fmt.Println(size)
	for g >= len(arr)/2 {
		if g-size[right] > len(arr)/2 {
			g -= size[right]
			res++
		} else if g-size[left] > len(arr) {
			g -= size[left]
			res++
		} else {
			return res
		}
	}
	return res
}
