package main

import (
	"strings"

	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
)

func main() {
	entityParser("and I quote: &quot;...&quot;")
}

func entityParser(text string) string {
	res := ""
	h := map[string]string{"&quot;": "\\\"", "&apos": "'", "&amp;": "&", "&gt;": ">",
		"&lt;": "<", "&frasl;": "/"}
	g := strings.Split(text, " ")

	for i, s := range g {
		if _, ok := h[s]; ok {
			res += h[s]
		} else {
			res += s
		}
		if i != len(g)-1 {
			res += " "
		}
	}
	fmt.Println(res)

	return res
}
