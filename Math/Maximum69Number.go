package main

func main() {

}

func maximum69Number(num int) int {
	temp, placeValue, last6Position := num, 1, 0

	for temp > 0 {
		if temp%10 == 6 {
			last6Position = placeValue
		}
		temp /= 10
		placeValue *= 10
	}
	return num + last6Position*3
}
