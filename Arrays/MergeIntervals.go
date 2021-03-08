package main

func main() {

}

func merge(intervals [][]int) [][]int {
	intervals = sortNotBuiltIn(intervals)
	h := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		for true {
			if len(h) == 0 {
				h = append(h, intervals[i])
				break
			}
			if h[len(h)-1][1] >= intervals[i][0] {
				if h[len(h)-1][0] >= intervals[i][0] {
					if h[len(h)-1][1] > intervals[i][1] {
						x := h[len(h)-1]
						h = h[:len(h)-1]
						h = append(h, []int{intervals[i][0], x[1]})
						break
					} else {
						h = h[:len(h)-1]
					}
				} else {
					x := h[len(h)-1]
					h = h[:len(h)-1]
					if x[1] > intervals[i][1] {
						h = append(h, x)
						break
					} else {
						h = append(h, []int{x[0], intervals[i][1]})
					}
					break
				}
			} else {
				h = append(h, intervals[i])
				break
			}
		}
	}
	return h
}

func sortNotBuiltIn(h [][]int) [][]int {
	for i := 1; i < len(h); i++ {
		if i >= 1 && (h[i-1][0] > h[i][0]) || (h[i-1][1] > h[i][1] && h[i-1][0] == h[i][0]) {
			h[i-1], h[i] = h[i], h[i-1]
			i -= 2
		}
	}
	return h
}
