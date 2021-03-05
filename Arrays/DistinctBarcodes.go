package main

func main() {

}

func rearrangeBarcodes(barcodes []int) []int {
	left := 1
	right := len(barcodes) - 2
	for left < right {
		if barcodes[right] != barcodes[left-1] && barcodes[left] != barcodes[right-1] {
			barcodes[right], barcodes[left] = barcodes[right], barcodes[left]
		}
	}
	return []int{}
}
