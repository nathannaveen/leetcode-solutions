package main

import "strconv"

func main() {

}

func validWordAbbreviation(word string, abbr string) bool {
	counter := 0
	for i := 0; i < len(word); i++ {
		stringNumber := ""
		if counter == len(abbr) {
			return false
		}
		for abbr[counter] <= 57 && abbr[counter] >= 48 && counter < len(abbr) {
			stringNumber += string(abbr[counter])
			counter++
		}
		if stringNumber != "" {
			g, _ := strconv.Atoi(stringNumber)
			i += g
		} else if string(abbr[counter]) != string(word[i]) {
			return false
		}
		counter++
	}
	return true
}
