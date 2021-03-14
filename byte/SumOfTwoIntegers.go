package main

func main() {

}

func getSum(a int, b int) int {
	res := 1
	res <<= b / 2
	if b%2 == 1 {
		res &= 1
	}
	b /= 2
	res <<= a / 2
	if a%2 == 1 {
		res &= 1
	}
	a /= 2
	return res
}
