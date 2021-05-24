package main

import "fmt"

func main() {

}

func accountsMerge(accounts [][]string) [][]string {
	m := make(map[string]string) // name, email
	res := [][]string{}

	for _, account := range accounts {
		shouldAddToRes := true

		for i := 1; i < len(account); i++ {
			fmt.Println(m, account[0], account[i])
			if m[account[0]] == account[i] {
				shouldAddToRes = false
				break
			} else {
				m[account[0]] = account[i]
			}
		}

		if shouldAddToRes {
			res = append(res, account)
		} else {
			for _, i := range res {
				if i[0] == account[0] {
					i = append(i, account[1:]...)
					break
				}
			}
		}
	}
	return res
}
