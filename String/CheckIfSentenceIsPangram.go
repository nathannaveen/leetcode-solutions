package main

func main() {

}

func checkIfPangram(sentence string) bool {
	arr, counter := make([]int, 26), 0

	for _, i := range sentence {
		arr[int(i-'a')]++
		if arr[int(i-'a')] == 1 {
			counter++
		}
	}

	return counter == 26
}
