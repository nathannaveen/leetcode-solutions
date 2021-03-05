package main

func concatenatedBinary(n int) int {
	res := 0
	for i := 1; i <= n; i++ {
		h := i
		g := []int{}
		for h > 0 {
			g = append(g, h&1)
			h >>= 1
		}
		g = reverseIntegersArray(g)
		for _, i2 := range g {
			res *= 2
			res += i2
		}
	}
	return res % 1000000007
}

func reverseIntegersArray(h []int) []int {
	left, right := 0, len(h)-1

	for left < right {
		h[left], h[right] = h[right], h[left]
		left++
		right--
	}

	return h
}
