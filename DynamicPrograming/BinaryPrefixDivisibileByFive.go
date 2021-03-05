package main

func main() {
}
func prefixesDivBy5(A []int) []bool {
	decimal := 0
	h := []bool{}
	for _, i2 := range A {
		decimal *= 2
		decimal += i2
		h = append(h, decimal%5 == 0)
	}
	return h
}
