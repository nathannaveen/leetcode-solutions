package main

func main() {

}

func checkValidString(s string) bool {
	numberOfExtra := 0
	stack := []rune{}

	for _, i := range s {
		if i == ')' {
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else if numberOfExtra == 0 {
				stack = append(stack, ')')
			} else if numberOfExtra > 0 {
				numberOfExtra--
			}
		} else if i == '(' {
			if len(stack) > 0 && stack[len(stack)-1] == ')' {
				stack = stack[:len(stack)-1]
			} else if numberOfExtra == 0 {
				stack = append(stack, '(')
			} else if numberOfExtra > 0 {
				numberOfExtra--
			}
		} else if i == '*' {
			numberOfExtra++
		}
	}
	return len(stack)-numberOfExtra <= 0
}
