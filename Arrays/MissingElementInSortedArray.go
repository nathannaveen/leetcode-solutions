package main

func main() {

}

func missingElement(nums []int, k int) int {
	h := []int{}
	for i := 1; i < 100000001; i++ {
		h = append(h, i)
	}
	counter := 0
	for _, num := range nums {
		h[num] = -1
	}
	for _, i := range h {
		if i != -1 {
			counter++
			if counter == k {
				return i
			}
		}
	}
	return -1
}
