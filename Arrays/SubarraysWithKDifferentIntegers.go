package main

func main() {

}

func subarraysWithKDistinct(A []int, K int) int {
	m := make(map[int]int)
	res := 0
	begining := 0

	for i := 0; i < len(A); i++ {
		m[A[i]]++
		for len(m) > K {
			m[A[begining]]--
			if m[A[begining]] == 0 {
				delete(m, A[begining])
			}
			begining++
		}
		if len(m) == K {
			res++
		}
	}

	return res
}
