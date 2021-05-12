package main

func main() {

}

func minNumberOperations(target []int) int {
	res := 0

	for i := 1; i < len(target); i++ {
		temp := target[i] - target[i-1]
		if temp > 0 {
			res += temp
		}
	}

	return res + target[0]
}
