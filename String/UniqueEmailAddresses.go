package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(numUniqueEmails([]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}))
}

func numUniqueEmails(emails []string) int {
	m := make(map[string]int)

	for _, email := range emails {
		end := strings.Index(email, "@")
		email = strings.ReplaceAll(email[:end], ".", "") + email[end:]
		start := strings.Index(email, "+")
		end = strings.Index(email, "@")
		if start != -1 {
			email = email[:start] + email[end:]
		}
		if m[email] == 0 {
			m[email] = 1
		}
	}

	return len(m)
}
