package main

func main() {

}

func arrangeCoins(n int) int {
	counter := 1
	for n > 0 {
		n -= counter
		if n == 0 {
			return counter
		}
		counter++
	}
	return counter - 2
}
