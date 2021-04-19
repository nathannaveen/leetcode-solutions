package main

import "fmt"

func main() {

}

func reverseWords(s []byte) {
	h := []byte{}
	end := len(s) - 1

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			fmt.Println(string(s[i : end+1]))
			h = append(h, s[i:end+1]...)
			end = i - 1
		}
	}
	h = append(h, ' ')
	h = append(h, s[0:end+1]...)
	fmt.Println(string(h))
	s = h
	fmt.Println(string(s))
}
