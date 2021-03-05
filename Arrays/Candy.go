package main

func main() {
}

func candy(ratings []int) int {
	h := make([]int, len(ratings))
	for i := range h {
		h[i] = 1
	}

	for i := 0; i < len(ratings)-1; i++ {
		if i == 0 && ratings[i] >= ratings[i+1] {
			h[i]++
		} else if i != 0 && (ratings[i] > ratings[i-1] || ratings[i] > ratings[i+1]) {
			h[i]++
		}
	}
	if ratings[len(ratings)-1] > ratings[len(ratings)-2] {
		h[len(ratings)-1]++
	}
	sum := 0
	for i := 0; i < len(h); i++ {
		sum += h[i]
	}
	return sum
}
