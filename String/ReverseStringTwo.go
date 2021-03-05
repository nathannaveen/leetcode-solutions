package main

func main() {
}

func reverseStr(s string, k int) string {
	if len(s) == 1 {
		return s
	}
	str := ""
	counter := 0
	Reverse := true
	for len(str) < len(s) {
		if Reverse {
			for i := counter + k - 1; i >= counter; i-- {
				if i == len(s) {
					return str
				}
				str += string(s[i])
			}
			counter += k
		} else {
			for i := counter; i < counter+k; i++ {
				if i == len(s) {
					return str
				}
				str += string(s[i])
			}
			counter += k
		}
		Reverse = !Reverse
	}
	if len(s)%k != 0 {
		str += string(s[len(s)-1])
	}
	return str
}
