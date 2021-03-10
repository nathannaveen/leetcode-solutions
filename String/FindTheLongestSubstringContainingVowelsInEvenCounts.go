package main

func main() {

}

func findTheLongestSubstring(s string) int {
	max := 0
	for i := 1; i < len(s); i++ {
		numberOfVowels := make([]int, 5)

		for j := i; j < len(s); j++ {
			// s[j : j + i]
			if j == i {
				for k := 0; k < i; k++ {
					numberOfVowels = whatVowelIsIt(s[k], numberOfVowels, 1)
				}
			} else {
				numberOfVowels = whatVowelIsIt(s[j-i-1], numberOfVowels, -1)
				numberOfVowels = whatVowelIsIt(s[j], numberOfVowels, 1)
			}
			if isEvenNumber(numberOfVowels) {
				max = i
				break
			}
		}
		//[0 3 2 2 0]
		//eleetminicowo XXXXXXXXXXXXXXXXXXXXXx
		//[0 3 2 3 0] l o
		//leetminicowor XXXXXXXXXXXXXXXXXXXXXx
		//[0 3 2 3 0] e e
		//eetminicoworo XXXXXXXXXXXXXXXXXXXXXx
		//[0 2 2 3 0] e p
		//etminicoworoe XXXXXXXXXXXXXXXXXXXXXx
		//13

		//[0 3 2 2 0]
		//eleetminicowo XXXXXXXXXXXXXXXXXXXXXx
		//[0 3 2 3 0] e o
		//leetminicowor XXXXXXXXXXXXXXXXXXXXXx
		//[0 3 2 3 0] l e
		//eetminicoworo XXXXXXXXXXXXXXXXXXXXXx
		//[0 2 2 3 0] e p
		//etminicoworoe XXXXXXXXXXXXXXXXXXXXXx
		//13
	}

	return max
}

func whatVowelIsIt(r uint8, h []int, number int) []int {
	switch r {
	case 'a':
		h[0] += number
	case 'e':
		h[1] += number
	case 'i':
		h[2] += number
	case 'o':
		h[3] += number
	case 'u':
		h[4] += number
	}
	return h
}

func isEvenNumber(h []int) bool {
	for _, i := range h {
		if i%2 == 1 {
			return false
		}
	}
	return true
}

//e
//e XXXXXXXXXXXXXXXXXXXXXx
//l XXXXXXXXXXXXXXXXXXXXXx
//l
//1
//e
//l
//el XXXXXXXXXXXXXXXXXXXXXx
//le XXXXXXXXXXXXXXXXXXXXXx
//le
//2
//e
//l
//e
//ele XXXXXXXXXXXXXXXXXXXXXx
//ele
//3
//e
//l
//e
//e
//elee XXXXXXXXXXXXXXXXXXXXXx
//leet XXXXXXXXXXXXXXXXXXXXXx
//eetm XXXXXXXXXXXXXXXXXXXXXx
//etmi XXXXXXXXXXXXXXXXXXXXXx
//tmin XXXXXXXXXXXXXXXXXXXXXx
//mini XXXXXXXXXXXXXXXXXXXXXx
//inic XXXXXXXXXXXXXXXXXXXXXx
//nico XXXXXXXXXXXXXXXXXXXXXx
//icow XXXXXXXXXXXXXXXXXXXXXx
//cowo XXXXXXXXXXXXXXXXXXXXXx
//owor XXXXXXXXXXXXXXXXXXXXXx
//woro XXXXXXXXXXXXXXXXXXXXXx
//woro
//4
//e
//l
//e
//e
//t
//eleet XXXXXXXXXXXXXXXXXXXXXx
//leetm XXXXXXXXXXXXXXXXXXXXXx
//eetmi XXXXXXXXXXXXXXXXXXXXXx
//etmin XXXXXXXXXXXXXXXXXXXXXx
//tmini XXXXXXXXXXXXXXXXXXXXXx
//minic XXXXXXXXXXXXXXXXXXXXXx
//inico XXXXXXXXXXXXXXXXXXXXXx
//nicow XXXXXXXXXXXXXXXXXXXXXx
//icowo XXXXXXXXXXXXXXXXXXXXXx
//cowor XXXXXXXXXXXXXXXXXXXXXx
//oworo XXXXXXXXXXXXXXXXXXXXXx
//oworo
//5
//e
//l
//e
//e
//t
//m
//eleetm XXXXXXXXXXXXXXXXXXXXXx
//leetmi XXXXXXXXXXXXXXXXXXXXXx
//eetmin XXXXXXXXXXXXXXXXXXXXXx
//etmini XXXXXXXXXXXXXXXXXXXXXx
//tminic XXXXXXXXXXXXXXXXXXXXXx
//minico XXXXXXXXXXXXXXXXXXXXXx
//inicow XXXXXXXXXXXXXXXXXXXXXx
//nicowo XXXXXXXXXXXXXXXXXXXXXx
//icowor XXXXXXXXXXXXXXXXXXXXXx
//coworo XXXXXXXXXXXXXXXXXXXXXx
//oworoe XXXXXXXXXXXXXXXXXXXXXx
//oworoe
//6
//e
//l
//e
//e
//t
//m
//i
//eleetmi XXXXXXXXXXXXXXXXXXXXXx
//leetmin XXXXXXXXXXXXXXXXXXXXXx
//eetmini XXXXXXXXXXXXXXXXXXXXXx
//eetmini
//7
//e
//l
//e
//e
//t
//m
//i
//n
//eleetmin XXXXXXXXXXXXXXXXXXXXXx
//leetmini XXXXXXXXXXXXXXXXXXXXXx
//eetminic XXXXXXXXXXXXXXXXXXXXXx
//etminico XXXXXXXXXXXXXXXXXXXXXx
//tminicow XXXXXXXXXXXXXXXXXXXXXx
//minicowo XXXXXXXXXXXXXXXXXXXXXx
//inicowor XXXXXXXXXXXXXXXXXXXXXx
//nicoworo XXXXXXXXXXXXXXXXXXXXXx
//icoworoe XXXXXXXXXXXXXXXXXXXXXx
//8
//e
//l
//e
//e
//t
//m
//i
//n
//i
//eleetmini XXXXXXXXXXXXXXXXXXXXXx
//leetminic XXXXXXXXXXXXXXXXXXXXXx
//eetminico XXXXXXXXXXXXXXXXXXXXXx
//etminicow XXXXXXXXXXXXXXXXXXXXXx
//tminicowo XXXXXXXXXXXXXXXXXXXXXx
//minicowor XXXXXXXXXXXXXXXXXXXXXx
//inicoworo XXXXXXXXXXXXXXXXXXXXXx
//nicoworoe XXXXXXXXXXXXXXXXXXXXXx
//9
//e
//l
//e
//e
//t
//m
//i
//n
//i
//c
//eleetminic XXXXXXXXXXXXXXXXXXXXXx
//leetminico XXXXXXXXXXXXXXXXXXXXXx
//eetminicow XXXXXXXXXXXXXXXXXXXXXx
//etminicowo XXXXXXXXXXXXXXXXXXXXXx
//tminicowor XXXXXXXXXXXXXXXXXXXXXx
//minicoworo XXXXXXXXXXXXXXXXXXXXXx
//minicoworo
//10
//e
//l
//e
//e
//t
//m
//i
//n
//i
//c
//o
//eleetminico XXXXXXXXXXXXXXXXXXXXXx
//leetminicow XXXXXXXXXXXXXXXXXXXXXx
//eetminicowo XXXXXXXXXXXXXXXXXXXXXx
//eetminicowo
//11
//e
//l
//e
//e
//t
//m
//i
//n
//i
//c
//o
//w
//eleetminicow XXXXXXXXXXXXXXXXXXXXXx
//leetminicowo XXXXXXXXXXXXXXXXXXXXXx
//eetminicowor XXXXXXXXXXXXXXXXXXXXXx
//eetminicowor
//12
//e
//l
//e
//e
//t
//m
//i
//n
//i
//c
//o
//w
//o
//eleetminicowo XXXXXXXXXXXXXXXXXXXXXx
//leetminicowor XXXXXXXXXXXXXXXXXXXXXx
//eetminicoworo XXXXXXXXXXXXXXXXXXXXXx
//etminicoworoe XXXXXXXXXXXXXXXXXXXXXx
//13
//e
//l
//e
//e
//t
//m
//i
//n
//i
//c
//o
//w
//o
//r
//eleetminicowor XXXXXXXXXXXXXXXXXXXXXx
//leetminicoworo XXXXXXXXXXXXXXXXXXXXXx
//leetminicoworo
//14
//e
//l
//e
//e
//t
//m
//i
//n
//i
//c
//o
//w
//o
//r
//o
//eleetminicoworo XXXXXXXXXXXXXXXXXXXXXx
//leetminicoworoe XXXXXXXXXXXXXXXXXXXXXx
//15
//e
//l
//e
//e
//t
//m
//i
//n
//i
//c
//o
//w
//o
//r
//o
//e
//eleetminicoworoe XXXXXXXXXXXXXXXXXXXXXx
//16
