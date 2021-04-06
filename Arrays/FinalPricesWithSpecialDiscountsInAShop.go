package main

func main() {

}

func finalPrices(prices []int) []int {
	minimumIndex := 0
	minimum := 1001

	for i, price := range prices {
		hasMinimum := true
		if minimumIndex <= i {
			hasMinimum = false
			for j := i + 1; j < len(prices); j++ {
				if prices[j] <= price {
					hasMinimum = true
					minimum = prices[j]
					break
				}
			}
		}

		if hasMinimum {
			prices[i] = price - minimum
		}
	}
	return prices
}
