package main

func customSortString(S string, T string) string {
	letters := make([]int, 26)
	res := ""
	for _, i := range T {
		letters[int(i-'a')]++
	}

	for _, i := range S {
		counter := letters[int(i-'a')]
		for j := 0; j < counter; j++ {
			res += string(i)
			letters[int(i-'a')]--
		}
	}
	for i := 0; i < 26; i++ {
		for j := 0; j < letters[i]; j++ {
			res += string(i + 'a')
		}
	}
	return res
}
