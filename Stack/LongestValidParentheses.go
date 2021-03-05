package main

func main() {

}

func longestValidParentheses(s string) int {
	h := []string{}
	counter := 0
	for _, i2 := range s {
		if len(h) != 0 && string(i2) == ")" {
			counter += 2
			h = h[:len(h)-1]
		} else if string(i2) == "(" {
			h = append(h, "(")
		}
	}
	return counter
}
