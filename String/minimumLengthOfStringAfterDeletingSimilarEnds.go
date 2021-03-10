package main

func minimumLength(s string) int {

	for len(s) > 1 && s[0] == s[len(s)-1] {
		end := s[len(s)-1]
		lenS := len(s)

		for i := 0; i < lenS; i++ {
			if s[len(s)-1] != s[0] {
				s = s[:len(s)-i-1]
				break
			}
		}
		lenS = len(s)
		for i := 0; i < lenS; i++ {
			if s[0] != end {
				s = s[0+i+1:]
				break
			}
		}
	}

	return len(s)
}
