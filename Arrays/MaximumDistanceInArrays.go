package main

func main() {

}

func maxDistance(arrays [][]int) int {
	maximum, minimum := -10001, 10001

	for _, array := range arrays {
		smallest, greatest := array[0], array[len(array)-1]

		if greatest > maximum && smallest < minimum {
			if greatest-maximum > minimum-smallest {
				maximum = greatest
			} else {
				minimum = smallest
			}
		} else if greatest > maximum {
			maximum = greatest
		} else if smallest < minimum {
			minimum = smallest
		}
	}
	return maximum - minimum
}
