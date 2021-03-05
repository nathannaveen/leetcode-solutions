package main

func getKth(lo int, hi int, k int) int {
	h := [][]int{}
	for i := lo; i <= hi; i++ {
		counter := 0
		n := i
		for n != 1 {
			counter++
			if n%2 == 0 {
				n /= 2
			} else {
				n = (n * 3) + 1
			}
		}
		g := []int{i, counter}
		h = append(h, g)
	}
	for i := 1; i < len(h); i++ {
		if (h[i-1][1] > h[i][1]) || (h[i-1][1] == h[i][1] && h[i-1][0] > h[i][0]) {
			g := h[i]
			for j := i - 1; j >= 0; j-- {
				if j >= 2 && h[j][1] > g[1] && h[j-1][1] < g[1] {
					h[j], h[j+1] = g, h[j]
				}
				if h[j][1] < g[1] {
					h[j] = g
					break
				} else if h[j][1] == g[1] && h[j][0] > g[0] {
					h[j] = g
					break
				}
				h[j+1] = h[j]
			}
		}
	}

	return h[k-1][0]
}
