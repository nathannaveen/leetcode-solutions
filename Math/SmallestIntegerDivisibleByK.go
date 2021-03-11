package main

func main() {

}

func smallestRepunitDivByK(K int) int {
	if (K%2 == 0 || K%5 == 0) && K != 1 {
		return -1
	}
	counter := 1
	integer := 1
	for {
		if integer%K == 0 {
			return counter
		}
		integer *= 10
		integer += 1
		counter++
	}
}
