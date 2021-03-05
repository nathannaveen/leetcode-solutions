package main

import "golang.org/x/tools/go/ssa/interp/testdata/src/fmt"

func main() {

}

func camelMatch(queries []string, pattern string) []bool {
	res := []bool{}
	patCounter := 0
	for _, query := range queries {
		remadePattern := ""
		for i := 0; i < len(query); i++ {
			if patCounter != len(pattern) && query[i] == pattern[patCounter] {
				remadePattern += string(query[i])
				patCounter++
			}
		}
		fmt.Println(remadePattern)
		res = append(res, remadePattern == pattern)
		fmt.Println(res)
	}
	return res
}
