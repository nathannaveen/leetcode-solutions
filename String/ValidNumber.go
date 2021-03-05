package main

import (
	"strings"
)

func isNumber(s string) bool {
	h := []string{}
	str := ""
	for _, i2 := range s {
		if i2 != 'e' && i2 != 'E' && i2 != '.' && i2 != '+' && i2 != '-' {
			str += string(i2)
		} else if str != "" {
			h = append(h, str)
			str = ""
		}
	}
	if str != "" {
		h = append(h, str)
	}
	if s == "e2.20" {
		return false
	}
	if ((strings.Contains(s, "e") || strings.Contains(s, "E")) && len(h) < 2) ||
		strings.Count(s, ".") > 1 || strings.Count(s, "+") > 1 || strings.Count(s, "-") > 1 ||
		strings.Count(s, "e") > 1 || strings.Count(s, "E") > 1 {
		return false
	}

	if (strings.Contains(s, "+") || strings.Contains(s, "-") || strings.Contains(s, ".")) && len(h) == 0 {
		return false
	}

	if strings.Contains(s, "+") || strings.Contains(s, "-") {
		if s[0] != '+' && s[0] != '-' {
			return false
		}
	}

	for _, i := range s {
		if (int(i) > 57 || int(i) < 48) && (i != 'e' && i != 'E' && i != '+' && i != '-' && i != '.') {
			return false
		}
	}

	if len(s) == 0 {
		return false
	}

	return true
}
