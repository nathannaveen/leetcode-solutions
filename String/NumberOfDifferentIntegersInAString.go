package main

import "fmt"

func main() {
	fmt.Println(numDifferentIntegers("leet1234code234"))
	fmt.Println(numDifferentIntegers("a1b01c001"))
}

func numDifferentIntegers(word string) int {
	digits := make(map[string]int)
	numbers := make(map[string]int)
	res := 0

	digits["0"]++
	digits["1"]++
	digits["2"]++
	digits["3"]++
	digits["4"]++
	digits["5"]++
	digits["6"]++
	digits["7"]++
	digits["8"]++
	digits["9"]++

	number := ""
	for i := 0; i < len(word); i++ {
		if digits[string(word[i])] == 1 {
			fmt.Println(number)
			if string(word[i]) == "0" && number == "" {
				fmt.Println(i)
				continue
			} else {
				fmt.Println(string(word[i]), "word")
				number += string(word[i])
			}
		} else {
			fmt.Println("taco")
			if number != "" && numbers[number] == 0 {
				numbers[number]++
				res++
				fmt.Println("added")
			}
			number = ""
		}
	}
	if number != "" && numbers[number] == 0 {
		res++
	}
	return res
}
