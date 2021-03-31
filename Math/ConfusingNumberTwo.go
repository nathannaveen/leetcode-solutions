package main

func main() {

}

func confusingNumberII(N int) int {
	counter := 0
	for i := 6; i < N; i++ {
		newNumber := 0
		b := true
		newN := i
		for newN > 0 {
			lastDigit := newN % 10
			newNumber *= 10
			switch lastDigit {
			case 0:
				newNumber += 0
			case 1:
				newNumber += 1
			case 6:
				newNumber += 9
			case 8:
				newNumber += 8
			case 9:
				newNumber += 6
			default:
				b = false
				break
			}

			newN /= 10
		}
		if newNumber != i && b == true {
			counter++
		}
	}
	return counter
}
