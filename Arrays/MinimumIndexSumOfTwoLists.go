package main

func main() {

}

func findRestaurant(list1 []string, list2 []string) []string {
	m := make(map[string]int)
	res := []string{}
	minimum := 2001

	for i, s := range list1 {
		m[s] = i
	}

	for i, s := range list2 {
		if _, ok := m[s]; ok {
			indexSum := m[s] + i
			if indexSum < minimum {
				minimum = indexSum
				res = []string{s}
			} else if indexSum == minimum {
				res = append(res, s)
			}
		}
	}
	return res
}
