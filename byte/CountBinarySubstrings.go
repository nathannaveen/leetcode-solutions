package main

func main() {

}

func countBinarySubstrings(s string) int {
	counter := 0
	for i := range s {
		smallCounter := 0
		for j := i; j < len(s); j++ {
			if string(s[j]) == "0" {
				smallCounter--
			} else {
				smallCounter++
			}

			if smallCounter == 0 {
				counter++
				break
			}
		}
	}
	return counter
}
