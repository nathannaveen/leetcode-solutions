package main

func main() {

}

func getNoZeroIntegers(n int) []int {
	for i := 1; i < n; i++ {
		containsZero := false
		a := i
		b := n - i
		for a > 0 {
			if a%10 == 0 {
				containsZero = true
				break
			}
			a /= 10
		}
		if !containsZero {
			for b > 0 {
				if b%10 == 0 {
					containsZero = true
					break
				}
				b /= 10
			}
		}
		if !containsZero {
			return []int{i, n - i}
		}
	}
	return []int{}
}
