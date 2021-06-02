package main

func main() {

}

func areSentencesSimilar(sentence1 []string, sentence2 []string, similarPairs [][]string) bool {
	if len(sentence1) != len(sentence2) {
		return false
	}
	length := len(sentence1)
	m := make(map[string]map[string]string)

	for _, pair := range similarPairs {
		if oldMap, ok := m[pair[0]]; !ok {
			newMap := make(map[string]string)
			newMap[pair[1]] = pair[1]
			m[pair[0]] = newMap
		} else {
			oldMap[pair[1]] = pair[1]
		}

		if oldMap, ok := m[pair[1]]; !ok {
			newMap := make(map[string]string)
			newMap[pair[0]] = pair[0]
			m[pair[1]] = newMap
		} else {
			oldMap[pair[0]] = pair[0]
		}
	}

	for i := 0; i < length; i++ {
		if oldMap, ok := m[sentence1[i]]; ok {
			if _, v := oldMap[sentence2[i]]; v {
				if sentence1[i] != sentence2[i] {
					return false
				}
			}
		}
	}

	return true
}

func arrayContains(h []string, n string) bool {
	for _, s := range h {
		if s == n {
			return true
		}
	}
	return false
}
