package main

func main() {

}

func minRemoveToMakeValid(s string) string {
	g := 0
	res := ""

	for _, i := range s {
		if i == '(' {
			g++
		} else if i == ')' {
			g--
		}
		if g >= 0 {
			res += string(i)
		}
	}
	if g > 0 {
		newRes := ""

		for _, i := range s {
			if i == '(' && g > 0 {
				g--
			} else {
				newRes += string(i)
			}
		}

		return newRes
	}
	return res
}
