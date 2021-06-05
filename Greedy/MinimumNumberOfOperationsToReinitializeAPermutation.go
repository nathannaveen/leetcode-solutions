package main

import "reflect"

func main() {

}

func reinitializePermutation(n int) int {
	perm := make([]int, n)
	arr := make([]int, n)
	res := 0

	for i := 0; i < n; i++ {
		perm[i] = i
		arr[i] = i
	}

	for true {
		for i := 0; i < len(arr); i++ {
			if i%2 == 1 {
				arr[i] = perm[n/2+(i-1)/2]
			} else {
				arr[i] = perm[i/2]
			}
		}
		if reflect.DeepEqual(perm, arr) {
			break
		}
		perm = arr
		res++
	}
	return res + 1
}
