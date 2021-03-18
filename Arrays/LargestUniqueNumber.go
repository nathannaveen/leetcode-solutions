package main

func main() {

}

func largestUniqueNumber(A []int) int {
	h := make(map[int]int)
	number := -1

	for _, i := range A {
		h[i]++
	}

	for i, i2 := range h {
		if i2 == 1 && i > number {
			number = i
		}
	}

	return number
}
