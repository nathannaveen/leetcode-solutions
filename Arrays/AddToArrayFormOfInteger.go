package main

func main() {
}

func addToArrayForm(A []int, K int) []int {
	if len(A) == 1 && A[0] == 0 && K == 0 {
		return []int{0}
	}
	h := []int{}
	next := 0
	for i := len(A) - 1; i >= 0; i-- {
		g := A[i] + next
		if K > 0 {
			g += K % 10
			K /= 10
		}
		next = 0
		if g > 9 {
			next = g / 10
		}
		l := g
		if g/10 >= 1 {
			l -= (g / 10) * 10
		}
		h = append(h, l) // problem
	}
	if next >= 1 {
		h = append(h, next)
	} else if K > 0 {
		for K > 0 {
			h = append(h, K%10)
			K /= 10
		}
	}
	return reverse(h)
}

func reverse(h []int) []int {
	left := 0
	right := len(h) - 1
	for left < right {
		h[left], h[right] = h[right], h[left]
		left++
		right--
	}
	return h
}
