package main

func main() {

}

func areAlmostEqual(s1 string, s2 string) bool {
	counter := 0
	str1Char, str2Char := uint8(' '), uint8(' ')

	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] && str1Char == ' ' {
			counter++
			str1Char = s1[i]
			str2Char = s2[i]
		} else if s1[i] != s2[i] && s2[i] != str1Char || s1[i] != str2Char {
			return false
		} else if s1[i] != s2[i] {
			counter++
		}
	}
	return counter == 0 || counter == 2
}
