package main

func main() {

}

func maxEnvelopes(envelopes [][]int) int {
	for i := 1; i < len(envelopes); i++ {
		if envelopes[i][0] <= envelopes[i-1][0] {
			placed := false
			for j := i - 1; j >= 0; j-- {
				if envelopes[j][0] < envelopes[i][0] {
					placed = true
					envelopes[j] = envelopes[i]
					break
				} else if envelopes[j][0] == envelopes[i][0] {
					if envelopes[j][1] < envelopes[i][1] {
						placed = true
						envelopes[j] = envelopes[i]
						break
					}
				}
				envelopes[j] = envelopes[j+1]
			}
			if !placed {
				envelopes[0] = envelopes[i]
			}

		}
	}
	counter := 0
	width := 0
	hight := 0
	for _, envelope := range envelopes {
		if width < envelope[0] && hight < envelope[1] {
			counter++
			width = envelope[0]
			hight = envelope[1]
		}
	}
	return counter
}
