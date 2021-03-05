package main

func main() {
}

func lastRemaining(n int) int {
	rightStart := false
	h := []int{}
	counter := n
	for i := 1; i <= n; i++ {
		h = append(h, i)
	}

	for counter > 1 {
		c := 0

		if rightStart == true {
			for i := len(h) - 1; i >= 0; i-- {
				if h[i] != -1 && counter == 1 {
					return h[i]
				}
				if h[i] != -1 {
					if c%2 == 0 {
						h[i] = -1
					}
					c++
				}
			}
			rightStart = false
		} else {
			for i := 0; i < len(h); i++ {
				if h[i] != -1 && counter == 1 {
					return h[i]
				}
				if h[i] != -1 {
					if c%2 == 0 {
						h[i] = -1
					}
					c++
				}
			}
			rightStart = true
		}
		counter /= 2
	}

	for i := 0; i < len(h); i++ {
		if h[i] != -1 {
			return h[i]
		}
	}

	return -1
}
