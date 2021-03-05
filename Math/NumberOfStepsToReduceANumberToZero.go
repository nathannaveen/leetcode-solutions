package main

func main() {

}

func numberOfSteps(num int) int {
	if num == 0 {
		return 0
	}
	counter := 0

	for num > 0 {
		if num%2 == 1 {
			counter++
		}
		counter++
		num /= 2
	}

	return counter - 1
}
