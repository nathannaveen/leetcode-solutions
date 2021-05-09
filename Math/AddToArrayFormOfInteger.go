package main

func addToArrayForm(num []int, k int) []int {
	numCounter := len(num) - 1
	carry := 0

	for k > 0 {
		n := k % 10
		g := n + num[numCounter] + carry

		if numCounter == -1 {
			num = append(num[:0], append([]int{0}, num[0:]...)...)
			numCounter++
		}

		if g > 9 {
			carry = g / 10
			num[numCounter] = g % 10
		} else {
			carry = 0
			num[numCounter] = g
		}
		numCounter--
		k /= 10
	}
	for carry != 0 {
		if numCounter == -1 {
			num = append(num[:0], append([]int{0}, num[0:]...)...)
			numCounter++
		}
		g := carry + num[numCounter]
		if g > 9 {
			carry = g / 10
			num[numCounter] = g % 10
		} else {
			carry = 0
			num[numCounter] = g
		}
	}
	return num
}
