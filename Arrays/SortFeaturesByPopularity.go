package main

import (
	"sort"
	"strings"
)

func main() {

}

func sortFeatures(features []string, responses []string) []string {
	m := make(map[string]int)

	for _, respons := range responses {
		h := strings.Split(respons, " ")
		sort.Strings(h)
		m[h[0]]++
		for i := 1; i < len(h); i++ {
			if h[i] != h[i-1] {
				m[h[i]]++
			}
		}
	}

	for i := 1; i < len(features); i++ {
		if i >= 1 && m[features[i-1]] < m[features[i]] {
			features[i-1], features[i] = features[i], features[i-1]
			i -= 2
		}
	}
	return features
}
