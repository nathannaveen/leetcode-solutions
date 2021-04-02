package main

func main() {
	//fmt.Println(maximumBeauty([]int{100, 1, 1, -3, 1}))
}

func maximumBeauty(flowers []int) int {
	m := make(map[int]int)
	beauty := -10001

	for _, flower := range flowers {
		m[flower]++
		if m[flower] >= 2 {
			isFlowers := false
			sum := 0
			oldSum := 0
			for _, flower2 := range flowers {
				if !isFlowers && flower2 == flower {
					sum += flower2
					isFlowers = true
				} else if isFlowers && flower2 == flower {
					sum += flower2
					if oldSum != 0 {
						oldSum = max(sum, oldSum)
					} else {
						oldSum = sum
					}
				} else if isFlowers && flower2 > 0 {
					sum += flower2
				}
			}
			beauty = max(beauty, oldSum)
		}
	}

	return beauty
}
